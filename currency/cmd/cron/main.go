package main

import (
	"context"
	"currencyService/currency/internal/clients/currency"
	"currencyService/currency/internal/config"
	"currencyService/currency/internal/db"
	"currencyService/currency/internal/repository"
	"currencyService/currency/internal/service"
	"currencyService/currency/internal/worker"
	"flag"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	configPath := flag.String("config", "./config", "path to the config file")

	flag.Parse()
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := db.InitConnection(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	currencyRepository := repository.NewCurrencyRepository(db)
	client := currency.NewCurrencyClient(cfg.ExternalAPI)
	currencyService := service.NewCurrencyService(currencyRepository, client)
	currencyWorker := worker.NewCurrencyWorker(cfg.Worker, currencyService)

	err = currencyWorker.Stat()
	if err != nil {
		log.Fatalf("error starting currency worker: %v", err)
	} else {
		log.Println("currency worker started successfully")
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()
	err = currencyWorker.Stop()
	if err != nil {
		log.Fatalf("error stopping currency worker: %v", err)
	} else {
		log.Println("currency worker stopped successfully")
	}
}
