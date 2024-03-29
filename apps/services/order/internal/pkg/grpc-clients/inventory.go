package grpcClients

import (
	"fmt"
	pbInventory "svetozar12/e-com/v2/api/v1/inventory/dist/proto"
	"svetozar12/e-com/v2/apps/services/order/internal/pkg/env"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var InventoryClient pbInventory.InventoryServiceClient

func initInventoryClient() {
	conn, err := grpc.Dial(env.Envs.INVENTORY_SERVICE_ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Errorf("Failed to dial bufnet: %v", err)
	}
	client := pbInventory.NewInventoryServiceClient(conn)
	InventoryClient = client
	fmt.Println("Inventory client READY")
}
