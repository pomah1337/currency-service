package handler

import (
	"currencyService/currency/internal/config"
	"currencyService/pkg/currency"
	"google.golang.org/grpc"
	"log"
	"net"
)

type CurrencyGrpcServer struct {
	s   *grpc.Server
	cfg config.GrpcCfg
	srv *CurrencyHandler
}

func NewGrpcServer(cfg config.GrpcCfg, srv *CurrencyHandler) *CurrencyGrpcServer {
	return &CurrencyGrpcServer{s: grpc.NewServer(), cfg: cfg, srv: srv}
}

func (c *CurrencyGrpcServer) StartServer() {
	listener, err := net.Listen("tcp", c.cfg.Host+":"+c.cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	currency.RegisterCurrencyServer(c.s, c.srv)
	go func() {
		if err := c.s.Serve(listener); err != nil {
			log.Fatalf("failed to start a grpc server: %v", err)
		}
	}()
}

func (c *CurrencyGrpcServer) StopServer() {
	c.s.Stop()
}
