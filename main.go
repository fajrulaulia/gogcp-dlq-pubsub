package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()

	log.Println("starting client")
	client, err := SetupClientGCPPubsub(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Wait message")

	var stop = make(chan bool)

	go func() {
		log.Println("Worker Default")
		subID := "mocca-topic-sub"
		err = client.Subscription(subID).Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
			log.Println("WOrker Defauilt got message", string(msg.Data))
			log.Println("WOrker Defauilt ordering key", string(msg.OrderingKey))
			if string(msg.Data) == "KIMINOTO" {
				msg.Nack()
			} else {
				msg.Ack()

			}
		})
		if err != nil {
			log.Default().Println("Error mocca-topic-sub", err.Error())
		}
	}()

	go func() {
		log.Println("Worker Error")
		subID := "mocca-dlq-topic-sub"
		err = client.Subscription(subID).Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
			log.Println("WOrker error  got message DLQ", string(msg.Data))
			log.Println("WOrker Defauilt  got key", string(msg.OrderingKey))
			msg.Ack()
		})
		if err != nil {
			log.Default().Println("Error mocca-dlq-topic-subb", err.Error())
		}
	}()

	<-stop

	log.Println("exit")

}

func SetupClientGCPPubsub(ctx context.Context) (*pubsub.Client, error) {
	projectID := "for-learning-363517"
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("pubsub.NewClient: %v", err)
	}

	return client, nil
}
