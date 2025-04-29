package worker

import (
	"context"
	"currencyService/currency/internal/config"
	"currencyService/currency/internal/service"
	"github.com/go-co-op/gocron/v2"
	"log"
)

type CurrencyWorker struct {
	currencyService service.CurrencyService
	scheduler       gocron.Scheduler
	cfg             *config.WorkerCfg
}

func NewCurrencyWorker(cfg *config.WorkerCfg, currencyService service.CurrencyService) *CurrencyWorker {
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err)
	}
	return &CurrencyWorker{
		currencyService: currencyService,
		scheduler:       s,
		cfg:             cfg,
	}
}

func (w CurrencyWorker) Stat() error {
	_, err := w.scheduler.NewJob(
		gocron.CronJob(w.cfg.Cron, false),
		//gocron.DurationJob(time.Second*10),
		gocron.NewTask(
			func() {
				err := w.currencyService.UpdateRate(
					context.Background(),
					w.cfg.Currencies.BaseCurrency,
					w.cfg.Currencies.TargetCurrency)
				if err != nil {
					log.Printf("failed to update rate: %v", err)
				} else {
					log.Println("rate updated successfully")
				}
			},
		),
	)
	if err != nil {
		return err
	}
	w.scheduler.Start()
	return nil
}

func (w CurrencyWorker) Stop() error {
	return w.scheduler.Shutdown()
}
