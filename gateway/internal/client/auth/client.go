package auth

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	pingPath          = "/ping"
	generateTokenPath = "/generate"
	validateTokenPath = "/validate"
	authHeader        = "Authorization"
)

type Client struct {
	baseURL    url.URL
	httpClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL:    url.URL{Scheme: "http", Host: baseURL},
		httpClient: &http.Client{},
	}
}

func (c *Client) Ping(ctx context.Context) (string, error) {
	reqUrl := c.baseURL.ResolveReference(&url.URL{Path: pingPath})
	req, err := http.NewRequestWithContext(ctx, "GET", reqUrl.String(), nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	return string(body), nil
}

func (c *Client) GenerateToken(ctx context.Context, login string) (string, error) {
	reqUrl := c.baseURL.ResolveReference(&url.URL{Path: generateTokenPath})
	q := reqUrl.Query()
	q.Set("login", login)
	reqUrl.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", reqUrl.String(), nil)
	if err != nil {
		return "", err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("failed to read response body: %w", err)
		}
		return string(body), nil
	case http.StatusBadRequest:
		return "", fmt.Errorf("bad request: failed to generate token for %s", login)
	default:
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
}

func (c *Client) ValidateToken(ctx context.Context, token string) error {
	reqUrl := c.baseURL.ResolveReference(&url.URL{Path: validateTokenPath})
	req, err := http.NewRequestWithContext(ctx, "GET", reqUrl.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Add(authHeader, "Bearer "+token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid token: %s", token)
	}
	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusBadRequest:
		return fmt.Errorf("bad request: token is not found in header")
	case http.StatusUnauthorized:
		return fmt.Errorf("invalid or expired token")
	default:
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
}
