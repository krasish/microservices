package main

import (
	"log"

	"github.com/krasish/microservices/order/config"
	"github.com/krasish/microservices/order/internal/adapters/db"
	"github.com/krasish/microservices/order/internal/adapters/grpc"
	"github.com/krasish/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}
	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}