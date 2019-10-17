package main

import (
	"log"
	"os"
	"time"

	"github.com/seguijoaquin/tennis-statistics/src/app/fetcher"

	"github.com/seguijoaquin/tennis-statistics/src/app/consumer"

	"github.com/seguijoaquin/tennis-statistics/src/app/producer"
)

func main() {
	log.Println("Starting app")
	switch os.Getenv("module") {
	case "producer":
		p, err := producer.NewProducer()
		for err != nil {
			time.Sleep(3 * time.Second)
			p, err = producer.NewProducer()
		}
		p.Run()
	case "consumer":
		c, err := consumer.NewConsumer()
		for err != nil {
			time.Sleep(3 * time.Second)
			c, err = consumer.NewConsumer()
		}
		c.Run()
	case "fetcher":
		f, err := fetcher.NewFetcher()
		for err != nil {
			time.Sleep(3 * time.Second)
			f, err = fetcher.NewFetcher()
		}
		f.Run()
	default:
		log.Fatalln("No module selected")
	}
}
