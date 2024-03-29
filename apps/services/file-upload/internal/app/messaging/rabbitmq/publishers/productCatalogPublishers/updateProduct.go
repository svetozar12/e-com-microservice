package productCatalogPublishers

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

func UpdateProductMessage(ch *amqp.Channel, productData map[string]interface{}) error {
	queueName := "product-update-queue"
	data, err := json.Marshal(productData)
	if err != nil {
		return err
	}
	err = ch.PublishWithContext(context.Background(),
		"",        // Exchange
		queueName, // Routing key
		false,     // Mandatory
		false,     // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)
	if err != nil {
		return err
	}
	return err
}
