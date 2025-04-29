package currency

import (
	"currencyService/pkg/currency"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewCurrencyGrpcClient(address string) (currency.CurrencyClient, *grpc.ClientConn, error) {
	clientConn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect: %w", err)
	}
	return currency.NewCurrencyClient(clientConn), clientConn, nil
}
