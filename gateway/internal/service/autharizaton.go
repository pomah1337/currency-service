package service

import (
	"context"
	"currencyService/gateway/internal/client/auth"
	"currencyService/gateway/internal/repository"
	"fmt"
)

type AuthService struct {
	repo   repository.UserRepository
	client auth.Client
}

func NewAuthService(repo repository.UserRepository, client auth.Client) *AuthService {
	return &AuthService{repo: repo, client: client}
}

func (a *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	user, err := a.repo.FindUser(username)
	if err != nil {
		return "", fmt.Errorf("user not found: %w", err)
	}
	if user.Password != password {
		return "", fmt.Errorf("wrong password")
	}

	token, err := a.client.GenerateToken(ctx, username)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	return token, nil
}

func (a *AuthService) Registration(ctx context.Context, username, password string) error {
	_, err := a.repo.FindUser(username)
	if err == nil {
		return fmt.Errorf("user already exists")
	}
	if err := a.repo.SaveUser(username, password); err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}
