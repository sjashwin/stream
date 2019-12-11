package stream

const (
	//Version of this package
	Version = "0.1.0"
)

// Stream interface
type Stream interface {
	Convert() <-chan int

	Sum() int
	Average() int
	Min() int
	Max() int
}

// Integer type pipeline
type Integer struct {
	Values []int
	len    int
}

// Convert given int slice to a stream
func (i *Integer) Convert() <-chan int {
	var out chan int = make(chan int)
	go func() {
		for _, num := range i.Values {
			out <- num
		}
		close(out)
	}()
	return out
}
