package service

import (
	"context"
	"currencyService/currency/internal/clients/currency"
	"currencyService/currency/internal/dto"
	"currencyService/currency/internal/repository"
	"fmt"
	"time"
)

type CurrencyService interface {
	UpdateRate(ctx context.Context, baseCurrency, targetCurrency string) error
	GetRate(cyt context.Context, baseCurrency, targetCurrency string) (*dto.ExchangeRate, error)
	GetHistory(ctx context.Context, base, target string, from, to time.Time) ([]dto.ExchangeRate, error)
}

type currencyService struct {
	repo   repository.CurrencyRepository
	client *currency.Client
}

func NewCurrencyService(repo repository.CurrencyRepository, client *currency.Client) CurrencyService {
	return &currencyService{repo: repo, client: client}
}

func (s *currencyService) UpdateRate(ctx context.Context, baseCurrency, targetCurrency string) error {
	rate, err := s.client.GetUsdRate()
	if err != nil {
		return fmt.Errorf("failed to get USD rate: %w", err)
	}

	err = s.repo.SaveRate(ctx,
		dto.ExchangeRate{BaseCurrency: baseCurrency, TargetCurrency: targetCurrency, Rate: rate, UpdateDate: time.Now()})

	if err != nil {
		return fmt.Errorf("failed to save USD rate: %w", err)
	}

	return nil
}

func (s *currencyService) GetRate(ctx context.Context, baseCurrency, targetCurrency string) (*dto.ExchangeRate, error) {
	rate, err := s.repo.GetRate(ctx, baseCurrency, targetCurrency)
	if err != nil {
		return nil, err
	}
	return rate, nil
}

func (s *currencyService) GetHistory(ctx context.Context, base, target string, from, to time.Time) ([]dto.ExchangeRate, error) {
	history, err := s.repo.GetHistory(ctx, base, target, from, to)
	if err != nil {
		return nil, err
	}
	return history, nil
}
