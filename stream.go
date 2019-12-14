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

// Integer type stream
type Integer struct {
	Values []int
	len    int
}

// Float type stream
type Float struct {
	Values []float64
	Len    int
}

// String type stream
type String struct {
	Values []string
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

// Convert a given float slice to a stream
func (f *Float) Convert() <-chan float64 {
	var out chan float64 = make(chan float64)
	go func() {
		for _, num := range f.Values {
			out <- num
		}
		close(out)
	}()
	return out
}

// Convert a give float slice to a string stream
func (s *String) Convert() <-chan string {
	var out chan string = make(chan string)
	go func() {
		for _, num := range s.Values {
			out <- num
		}
		close(out)
	}()
	return out
}
