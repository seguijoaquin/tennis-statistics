package main

import (
	"log"
	"os"

	"github.com/seguijoaquin/tennis-statistics/src/app/consumer"

	"github.com/seguijoaquin/tennis-statistics/src/app/producer"
)

func main() {
	log.Println("Starting app")
	switch os.Getenv("module") {
	case "producer":
		p := producer.NewProducer()
		p.Run()
	case "consumer":
		c := consumer.NewConsumer()
		c.Run()
	default:
		log.Fatalln("No module selected")
	}
}
