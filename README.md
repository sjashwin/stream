## Stream Processing API for Golang

This repository is a stream processing api for golang. The api uses channels to communicate between
stream. This api is meant to be light weight to provide developers with basic stream solutions.

### Using the API

First, need to import the stream api package and use the `stream.Integer` to convert a slice of ints to stream.

```go
package main

import "github.com/sjashwin/stream"

func main(){
    var data stream.Integer = stream.Integer{
        Values: []int{1, 2, 3, 4, 5},
    }
}
```
#### How to use it

You can use any method that implements the Stream interface. Yes, you are right `stream.Integer` implements the `Stream` interface. To get the sum of all the elements in the given slice, simple:

```go
package main

import (
    "fmt"
    "github.com/sjashwin/stream"
)

func main(){
    var data stream.Integer = stream.Integer{
        Values: []int{1, 2, 3, 4, 5},
    }

    fmt.Println("Sum of all elements:", data.Sum()) // Output: 15
}
```

Apply a filter on the stream

```go
import (
    "fmt"
    "github.com/sjashwin/stream"
)

func main(){
    var data stream.Integer = stream.Integer{
        Values: []int{1, 2, 3, 4, 5},
    }

    var even func(num int) bool

    even = func(num int) bool{
        return num % 2 == 0
    }

    data.Filter(even)
    fmt.Println(data.Values) // [2, 4]
}
```

Map a slice by passing a desired method

```go
import (
    "fmt"
    "github.com/sjashwin/stream"
)

func main(){
    var data stream.Integer = stream.Integer{
        Values: []int{1, 2, 3, 4, 5},
    }

    var square(num int) int
    
    square = func(num int) int{
        return num*num
    }

    data.Map(square)
    fmt.Println(data.Values) // [1, 4, 9, 16, 25]
}
```