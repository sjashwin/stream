package stream

import (
	"errors"
	"reflect"
)

const (
	//Version of this package
	Version = "0.1.0"
)

// Stream interface
type Stream interface {
	Convert() <-chan interface{}

	Sum() interface{}
	Average() interface{}
	Min() interface{}
	Max() interface{}
	Len() int
	Val() interface{}
}

// Integer type stream
type Integer []int

// Float type stream
type Float []float64

// Convert given int slice to a stream
func (i *Integer) Convert() <-chan interface{} {
	var out chan interface{} = make(chan interface{})
	go func() {
		for _, num := range *i {
			out <- num
		}
		close(out)
	}()
	return out
}

// Convert a given float slice to a stream
func (f *Float) Convert() <-chan interface{} {
	var out chan interface{} = make(chan interface{})
	go func() {
		for _, num := range *f {
			out <- num
		}
		close(out)
	}()
	return out
}

// Add two stream objects
func Add(s1, s2 Stream) (Stream, error) {
	var output Stream
	if reflect.TypeOf(s1) != reflect.TypeOf(s2) {
		return nil, errors.New("error: cannnot do add operation on distinct type")
	}

	switch s1.(type) {
	case *Integer:
		output = addInteger(s1.(*Integer), s2.(*Integer))
	case *Float:
		output = addFloat(s1.(*Float), s2.(*Float))
	default:
		return nil, errors.New("error: unsupported type")
	}
	return output, nil
}

func addInteger(num1, num2 *Integer) Stream {
	if num1.Len() != num2.Len() {
		padding(num1, num2)
	}
	var output []int = make([]int, num1.Len())
	for i := 0; i < num1.Len(); i++ {
		output[i] = (*num1)[i] + (*num2)[i]
	}
	var out Integer = output
	return &out
}

func addFloat(num1, num2 *Float) Stream {
	if num1.Len() != num2.Len() {
		padding(num1, num2)
	}
	var output []float64 = make([]float64, num1.Len())
	for i := 0; i < num1.Len(); i++ {
		output[i] = (*num1)[i] + (*num2)[i]
	}
	var out Float = output
	return &out
}

func padding(num1, num2 Stream) error {
	var temp interface{}
	switch num1.(type) {
	case *Integer:
		if num1.Len() > num2.Len() {
			temp = make([]int, num1.Len()-num2.Len())
			*num2.(*Integer) = append(*num2.(*Integer), temp.([]int)...)
		} else {
			temp = make([]int, num1.Len()-num2.Len())
			*num1.(*Integer) = append(*num1.(*Integer), temp.([]int)...)
		}
	case *Float:
		if num1.Len() > num2.Len() {
			temp = make([]float64, num1.Len()-num2.Len())
			*num2.(*Float) = append(*num2.(*Float), temp.([]float64)...)
		} else {
			temp = make([]float64, num1.Len()-num2.Len())
			*num1.(*Float) = append(*num1.(*Float), temp.([]float64)...)
		}
	default:
		return errors.New("error: unsupported type")
	}
	return nil
}
