package dto

import (
	"currencyService/pkg/currency"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type ExchangeRate struct {
	ID             int       `db:"id"`
	BaseCurrency   string    `db:"base_currency"`
	TargetCurrency string    `db:"target_currency"`
	Rate           float64   `db:"rate"`
	UpdateDate     time.Time `db:"update_date"`
}

func ExchangeRateToGrpcRate(r *ExchangeRate) *currency.Rate {
	return &currency.Rate{
		Date: timestamppb.New(r.UpdateDate),
		Rate: r.Rate,
	}
}
