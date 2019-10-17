package fetcher

import "log"

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
