package filter

// Filter represents an object that consumes messages from a source
type Filter struct {
	source         string
	destination    string
	filterFunction Function
}

// Function interface used to execute filtering operations
type Function func([]byte) ([]byte, error)

// NewFilter returns a Filter object
func NewFilter(
	src, dst string,
	function Function,
) (*Filter, error) {
	return &Filter{
		source:         src,
		destination:    dst,
		filterFunction: function,
	}, nil
}

// Run starts the Filter object
func (p *Filter) Run() {
	// source, err := communication.GetBroker(p.source)
	// utils.FailOnError(err, "")
	// defer source.Close()

	// dest, err := communication.GetBroker(p.destination)
	// utils.FailOnError(err, "")
	// defer dest.Close()

	// // TODO: Exchange things...
	// msgChannel := source.Consume()
	// for msg := range msgChannel {
	// 	output, err := p.filterFunction(msg)
	// 	// TODO: Check for errors
	// 	dest.Send(output)
	// 	output.Ack(False)
	// }
}
