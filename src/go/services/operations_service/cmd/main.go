package main

import (
	"context"

	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"truenorth/packages/logger"
	"truenorth/packages/pubsub/consumer"
	"truenorth/packages/service_conf"
	operationspb "truenorth/pb/operations"
	"truenorth/services/operations_service/config"
	"truenorth/services/operations_service/consumers"
	controller "truenorth/services/operations_service/controllers"
)

var log = logger.GetLogger()

func main() {
	err := config.Load()
	ctx := context.TODO()
	if err != nil {
		log.Fatal(ctx, "cannot load config: ", err)
	}

	listener, s := service_conf.InitServer(
		service_conf.WithDatabase(),
		service_conf.WithMigrations(),
		service_conf.WithPort(config.Config.Port),
	)
	sqsConsumer := consumer.NewConsumer()
	sqsConsumer.AddQueue(config.Config.UserCreatedQueue, consumers.UserCreatedController)
	sqsConsumer.Start()

	localServer := controller.NewServer()
	healthgrpc.RegisterHealthServer(s, localServer)
	operationspb.RegisterOperationsServiceServer(s, localServer)
	reflection.Register(s)
	log.Infow(ctx, "Serving gRPC", "PORT: ", config.Config.Port)

	if err := s.Serve(listener); err != nil {
		log.Fatalf(ctx, "failed to serve: %v", err)
	}
}
