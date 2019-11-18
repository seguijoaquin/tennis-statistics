package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/seguijoaquin/tennis-statistics/common/communication/broker"

	"github.com/seguijoaquin/tennis-statistics/common/domain"

	"github.com/seguijoaquin/tennis-statistics/common/utils"
)

// Producer represents an object that produces messages
type Producer struct {
	broker     *broker.MessageBroker
	feedChan   chan []string
	doneChan   chan bool
	doneFilter chan bool
	args       []string
}

// NewProducer returns a Producer object
func NewProducer(args []string) (*Producer, error) {
	feedChan := make(chan []string)
	doneChan := make(chan bool)
	doneFilter := make(chan bool)

	broker, _ := broker.GetBroker(args[1])

	return &Producer{
		broker:     broker,
		feedChan:   feedChan,
		doneChan:   doneChan,
		doneFilter: doneFilter,
		args:       args}, nil
}

func (p *Producer) getExchangeName() string {
	return p.args[1]
}

func (p *Producer) getKey() string {
	return p.args[2]
}

func (p *Producer) buildBody(record []string) []byte {
	//TODO: build body from any source
	body, err := json.Marshal(domain.BuildGameFeedDTO(record))
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func (p *Producer) buildEndMessage() []byte {
	body, err := json.Marshal(domain.BuildEndMessage())
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func (p *Producer) notifyEnd() {
	log.Printf("Notify end...\n")
	p.broker.Post(
		p.getExchangeName(),
		p.getKey(),
		p.buildEndMessage())
	p.doneFilter <- true
}

func (p *Producer) dataFilter() {
	for {
		select {
		case feed := <-p.feedChan:
			body := p.buildBody(feed)
			err := p.broker.Post(
				p.getExchangeName(),
				p.getKey(),
				body,
			)
			utils.FailOnError(err, "Failed to publish a message")
		case <-p.doneChan:
			log.Printf("Finish Publishing...\n")
			goto end_loop
		}
	}
end_loop:
	p.notifyEnd()
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
func (p *Producer) Run() {
	log.Println("Producing...")

	go p.dataFilter()

	err := filepath.Walk("/data", p.walkFunction)
	if err != nil {
		fmt.Printf("error walking the path: %v\n", err)
		return
	}

	p.doneChan <- true
	<-p.doneFilter
}

// Close releases all resources taken by a Producer
func (p *Producer) Close() {
	p.broker.Close()
}

func main() {
	utils.WaitForDependencies()
	args := os.Args
	p, err := NewProducer(args)
	for err != nil {
		time.Sleep(3 * time.Second)
		p, err = NewProducer(args)
	}
	p.Run()
	p.Close()
}
