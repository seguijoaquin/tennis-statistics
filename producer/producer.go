package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/seguijoaquin/tennis-statistics/common"
)

// Producer represents an object that produces messages
type Producer struct {
	amqpChan   *amqp.Channel
	amqpConn   *amqp.Connection
	feedChan   chan []string
	doneChan   chan bool
	args []string
}

// NewProducer returns a Producer object
func NewProducer(args []string) (*Producer, error) {
	// TODO: Try to initialize dependencies (rabbit)
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	err = ch.ExchangeDeclare(
		args[1], // name
		"topic",           // type
		true,              // durable
		false,             // auto-deleted
		false,             // internal
		false,             // no-wait
		nil,               // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	feedChan := make(chan []string)
	doneChan := make(chan bool)

	return &Producer{
		amqpChan:   ch,
		amqpConn:   conn,
		feedChan:   feedChan,
		doneChan:   doneChan,
		args: args}, nil
}

func (p *Producer) getExchangeName() string {
	return p.args[1]
}

func (p *Producer) getRoutingKey() string {
	return p.args[2]
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (p *Producer) buildBody(record []string) []byte {
	body, err := json.Marshal(common.BuildGameFeedDTO(record))
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func (p *Producer) sendFeed(wg *sync.WaitGroup) {
	for {
		select {
		case feed := <-p.feedChan:
			body := p.buildBody(feed)
			err := p.amqpChan.Publish(
				p.getExchangeName(), // exchange
				p.getRoutingKey(),      // routing key
				false,             // mandatory
				false,             // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        body,
				})
			failOnError(err, "Failed to publish a message")
		case <-p.doneChan:
			log.Printf("Finish Publishing...\n")
			break
		}
	}

	wg.Done()
}

func (p *Producer) walkFunction(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
		return err
	}
	fmt.Printf("visited file or dir: %q\n", path)

	if info.IsDir() || filepath.Ext(info.Name()) != ".csv" {
		fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
	} else {
		f, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		r := csv.NewReader(f)

		// Let's assume the file has a valid headers first line
		_, err = r.Read()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatal(err)
		}

		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			p.feedChan <- record
		}
	}
	return nil
}

// Run starts the producer object
func (p *Producer) Run(wg *sync.WaitGroup) {
	log.Println("Producing...")

	go p.sendFeed(wg)

	err := filepath.Walk("/data", p.walkFunction)
	if err != nil {
		fmt.Printf("error walking the path: %v\n", err)
		return
	}

	p.doneChan <- true
}

// Close releases all resources taken by a Producer
func (p *Producer) Close() {
	p.amqpChan.Close()
	p.amqpConn.Close()
}

func main() {
	common.WaitForDependencies()
	args := os.Args
	p, err := NewProducer(args)
	var wg sync.WaitGroup
	for err != nil {
		time.Sleep(3 * time.Second)
		p, err = NewProducer(args)
	}
	p.Run(&wg)
	wg.Wait()
	p.Close()
}
