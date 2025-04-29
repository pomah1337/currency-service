package currency

import (
	"currencyService/currency/internal/config"
	"currencyService/currency/internal/dto"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
)

type Client struct {
	url string
}

func NewCurrencyClient(cfg config.ExternalApiCfg) *Client {
	return &Client{url: cfg.Url}
}

func (c *Client) GetUsdRate() (float64, error) {

	resp, err := http.Get(c.url)
	if err != nil {
		return math.NaN(), fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return math.NaN(), fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return math.NaN(), fmt.Errorf("failed to read response body: %w", err)
	}

	data := dto.ResponseData{}
	json.Unmarshal(body, &data)

	if data.Rub == nil {
		return math.NaN(), fmt.Errorf("missing 'rub' in response")
	}
	if _, ok := data.Rub["usd"]; !ok {
		return math.NaN(), fmt.Errorf("missing 'usd' in response")
	}
	return data.Rub["usd"], nil
}
