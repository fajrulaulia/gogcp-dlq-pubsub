package main

import (
	"context"
	"fmt"
	"log"
	"strings"

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

	go func(stop chan bool) {
		log.Println("----- Worker Default -----")
		subID := "mocca-topic-sub"
		err = client.Subscription(subID).Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
			log.Println("Worker Default [Message]", string(msg.Data))
			log.Println("Worker Default [Ordering Key]", string(msg.OrderingKey))
			if strings.ToUpper(string(msg.Data)) == "KIMINOTO" || strings.ToUpper(string(msg.Data)) == "EXIT" {
				if strings.ToUpper(string(msg.Data)) == "EXIT" {
					stop <- true
				}
				msg.Nack()
			} else {
				msg.Ack()

			}
		})
		if err != nil {
			log.Println("Worker Default Error", err.Error())
		}
	}(stop)

	go func() {
		log.Println("----- Worker DLQ -----")
		subID := "mocca-dlq-topic-sub"
		err = client.Subscription(subID).Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
			log.Println("Worker DLQ [Message]", string(msg.Data))
			log.Println("Worker DLQ [Ordering Key]", string(msg.OrderingKey))
			msg.Ack()
		})
		if err != nil {
			log.Println("Worker DLQ Error", err.Error())
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
