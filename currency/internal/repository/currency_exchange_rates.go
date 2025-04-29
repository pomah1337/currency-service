package repository

import (
	"context"
	"currencyService/currency/internal/dto"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

var ErrNotFound = errors.New("not found")

type CurrencyRepository interface {
	SaveRate(ctx context.Context, rate dto.ExchangeRate) error
	GetRate(ctx context.Context, base, target string) (*dto.ExchangeRate, error)
	GetHistory(ctx context.Context, base, target string, from, to time.Time) ([]dto.ExchangeRate, error)
}
type currencyRepo struct {
	db *sqlx.DB
}

func NewCurrencyRepository(db *sqlx.DB) CurrencyRepository {
	return &currencyRepo{db: db}
}

func (r *currencyRepo) SaveRate(ctx context.Context, rate dto.ExchangeRate) error {
	query := `
    INSERT INTO currency_exchange_rates (base_currency, target_currency, rate, update_date)
    VALUES (:base_currency, :target_currency, :rate, :update_date);
    `

	_, err := r.db.NamedExecContext(ctx, query, rate)
	return err
}

func (r *currencyRepo) GetRate(ctx context.Context, base, target string) (*dto.ExchangeRate, error) {
	var rate dto.ExchangeRate
	query := `
    SELECT id, base_currency, target_currency, rate, update_date
    FROM currency_exchange_rates
    WHERE base_currency = $1 AND target_currency = $2
    ORDER BY update_date DESC;
    `

	err := r.db.GetContext(ctx, &rate, query, base, target)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	return &rate, nil
}

func (r *currencyRepo) GetHistory(ctx context.Context, base, target string, from, to time.Time) ([]dto.ExchangeRate, error) {
	var rates []dto.ExchangeRate
	query := `
    SELECT id, base_currency, target_currency, rate, update_date
    FROM currency_exchange_rates
    WHERE base_currency = $1 AND target_currency = $2 AND update_date BETWEEN $3 AND $4
    ORDER BY update_date;
    `
	err := r.db.SelectContext(ctx, &rates, query, base, target, from, to)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	return rates, nil
}
