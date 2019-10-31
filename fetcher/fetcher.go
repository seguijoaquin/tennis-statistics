package main

import (
	"log"
	"time"

	"github.com/seguijoaquin/tennis-statistics/common"
)

// Fetcher represents an object that is responsible for reading csv file lines
// and deliver them to a feed queue
type Fetcher struct {
}

// NewFetcher returns a Fetcher object
func NewFetcher() (*Fetcher, error) {
	// TODO: try to instantiate dependencies before creating object (rabbit)
	return &Fetcher{}, nil
}

// Run starts the fetcher object
func (p *Fetcher) Run() {
	log.Println("Fetching file...")
}

func main() {
	common.WaitForDependencies()
	f, err := NewFetcher()
	for err != nil {
		time.Sleep(3 * time.Second)
		f, err = NewFetcher()
	}
	f.Run()
}
