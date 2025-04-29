package main

import (
	"context"
	"currencyService/gateway/internal/client/auth"
)

func main() {
	authClient := auth.NewClient("localhost:8082")
	_, err := authClient.Ping(context.Background())
	if err != nil {
		panic(err)
	}

}
