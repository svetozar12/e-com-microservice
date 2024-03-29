package bootstrap

import (
	"log"
	"net"
	"svetozar12/e-com/v2/apps/services/product-catalog/internal/app/databases/postgres"
	"svetozar12/e-com/v2/apps/services/product-catalog/internal/app/messaging/rabbitmq"
	"svetozar12/e-com/v2/apps/services/product-catalog/internal/app/messaging/rabbitmq/consumers/productCatalogConsumers"
	"svetozar12/e-com/v2/apps/services/product-catalog/internal/app/services/product"
	"svetozar12/e-com/v2/apps/services/product-catalog/internal/pkg/constants"
	"svetozar12/e-com/v2/apps/services/product-catalog/internal/pkg/env"

	"google.golang.org/grpc"
)

func Bootstrap() {
	env.InitConfig()
	instance, err := rabbitmq.GetRabbitMQInstance(env.Envs.RABBIT_MQ_CONNECTION_STRING)
	if err != nil {
		panic(err)
	}
	ch, err := instance.CreateChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	defer instance.Close()

	instance.DeclareQueues(constants.ProductUpdateQueueName,
		constants.FileDeleteQueueName,
		constants.FileUploadQueueName)

	go productCatalogConsumers.ConsumeProductUpdateMessage(ch)

	grpcAddr := ":" + env.Envs.Port
	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		panic(err)
	}
	postgres.InitPostgres()
	s := grpc.NewServer()
	product.InitProductService(s)
	println("Product Catalog service started on port", grpcAddr)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
