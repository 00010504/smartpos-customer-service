package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/Invan2/invan_customer_service/config"
	"github.com/Invan2/invan_customer_service/pkg/logger"
	"google.golang.org/grpc"
)

func main() {

	cfg := config.Load()

	log := logger.New(cfg.LogLevel, cfg.ServiceName)

	server := grpc.NewServer()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.HttpPort))
	if err != nil {
		log.Error("failed to listen: %v", logger.Error(err))
		return
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		server.GracefulStop()
	}()

	if err := server.Serve(lis); err != nil {
		log.Error("error", logger.Error(err))
		return
	}
}
