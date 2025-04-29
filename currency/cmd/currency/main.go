package main

import (
	"context"
	"currencyService/currency/internal/clients/currency"
	"currencyService/currency/internal/config"
	"currencyService/currency/internal/db"
	"currencyService/currency/internal/handler"
	"currencyService/currency/internal/repository"
	"currencyService/currency/internal/service"
	"flag"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	configPath := flag.String("config", "currency/internal/config/config.example.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := db.InitConnection(cfg.DB)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	repo := repository.NewCurrencyRepository(db)
	client := currency.NewCurrencyClient(cfg.ExternalAPI)
	svc := service.NewCurrencyService(repo, client)
	hnd := handler.NewCurrencyHandler(svc)
	grpcServer := handler.NewGrpcServer(cfg.Grpc, hnd)

	grpcServer.StartServer()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()

	grpcServer.StopServer()
}
