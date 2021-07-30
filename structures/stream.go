package structures

const (
	//Version of this package
	Version = "0.1.1"
)

// Stream interface
type Stream interface {
	Integer
	Integer8
	Integer32
	Integer64
	Byte
	Rune
}

type Integer interface {
	Sum() int
	Average() int
	Min() int
	Max() int
}

type Integer8 interface {
	Sum() int
	Average() int
	Min() int
	Max() int
}

type Integer16 interface {
	Sum() int
	Average() int
	Min() int
	Max() int
}

type Integer32 interface {
	Sum() int
	Average() int
	Min() int
	Max() int
}

type Integer64 interface {
	Sum() int
	Average() int
	Min() int
	Max() int
}

type Byte interface {
	Sum() int
	Average() int
	Min() int
	Max() int
}

type Rune interface {
	Sum() int
	Average() int
	Min() int
	Max() int
}
