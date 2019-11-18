package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/seguijoaquin/tennis-statistics/common/communication/broker"
	"github.com/seguijoaquin/tennis-statistics/common/domain"
	"github.com/seguijoaquin/tennis-statistics/common/utils"
)

// Consumer represents an object that consumes messages
type Consumer struct {
}

// NewConsumer returns a Consumer object
func NewConsumer() (*Consumer, error) {
	return &Consumer{}, nil
}

// Run starts the consumer object
func (p *Consumer) Run() {
	log.Println("Consuming...")

	broker, err := broker.GetBroker("feeds_topic")
	utils.FailOnError(err, "Error getting broker")
	defer broker.Close()

	broker.DeclareUnnamedMessageQueue()
	broker.BindTopicsToQueue("feeds_topic", os.Args[1:])

	msgs, err := broker.GetMessages()
	utils.FailOnError(err, "Failed to register a consumer")

	done := make(chan bool)

	go func() {
		for d := range msgs {
			var dto domain.GameFeedDTO
			json.Unmarshal(d.Body, &dto)

			log.Printf(" [x] %s", d.Body)

			if dto.Finished {
				break
			}
		}
		log.Printf("Finished consuming messages")
		done <- true
	}()

	log.Printf(" [*] Waiting for Feed ...")
	<-done
}

func main() {
	utils.WaitForDependencies()
	c, _ := NewConsumer()
	c.Run()
}
