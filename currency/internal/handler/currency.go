package handler

import (
	"context"
	"currencyService/currency/internal/dto"
	"currencyService/currency/internal/service"
	"currencyService/currency/internal/tools"
	"currencyService/pkg/currency"
	"log"
)

type CurrencyHandler struct {
	currency.CurrencyServer
	service service.CurrencyService
}

func NewCurrencyHandler(service service.CurrencyService) *CurrencyHandler {
	return &CurrencyHandler{service: service}
}

func (c *CurrencyHandler) GetRate(ctx context.Context, request *currency.GetRateRequest) (*currency.GetRateResponse, error) {
	log.Println("GetRate")
	rate, err := c.service.GetRate(ctx, request.GetBaseCurrency(), request.GetTargetCurrency())
	if err != nil {
		log.Println(err)
		return nil, tools.GrpcError(err)
	}
	log.Println("success response")
	return &currency.GetRateResponse{Rate: dto.ExchangeRateToGrpcRate(rate)}, nil
}

func (c *CurrencyHandler) GetHistory(ctx context.Context, request *currency.GetHistoryRequest) (*currency.GetHistoryResponse, error) {
	log.Println("GetHistory")
	history, err := c.service.GetHistory(ctx,
		request.GetBaseCurrency(),
		request.GetTargetCurrency(),
		request.GetStartDate().AsTime(),
		request.GetEndDate().AsTime())
	if err != nil {
		log.Println(err)
		return nil, tools.GrpcError(err)
	}
	rates := make([]*currency.Rate, len(history))
	for _, rate := range history {
		rates = append(rates, dto.ExchangeRateToGrpcRate(&rate))
	}
	log.Println("success response")
	return &currency.GetHistoryResponse{Rates: rates}, nil

}
