package stream

import "sort"

// Sum all the values in a given stream
func (i *Integer) Sum() interface{} {
	var sum int
	for num := range i.Convert() {
		sum += num.(int)
	}
	return sum
}

// Len of an Integer stream type
func (i *Integer) Len() int {
	return len(*i)
}

// Len of a Float stream type
func (f *Float) Len() int {
	return len(*f)
}

// Val stored in a stream
func (i *Integer) Val() interface{} {
	return *i
}

// Val stored in a stream
func (f *Float) Val() interface{} {
	return *f
}

// Average of all values in a given stream
func (i *Integer) Average() interface{} {
	var sum int
	for num := range i.Convert() {
		sum += num.(int)
	}
	return sum / i.Len()
}

// Min of all values in a given stream
func (i *Integer) Min() interface{} {
	min := <-i.Convert()
	for num := range i.Convert() {
		if num.(int) < min.(int) {
			min = num
		}
	}
	return min
}

//Max of all the values in a given stream
func (i *Integer) Max() interface{} {
	max := <-i.Convert()
	for num := range i.Convert() {
		if num.(int) > max.(int) {
			max = num
		}
	}
	return max
}

// Filter values from a given integer slice
func (i *Integer) Filter(f func(int) bool) {
	var buffer Integer = make(Integer, 0, i.Len())
	for num := range i.Convert() {
		if f(num.(int)) {
			buffer = append(buffer, num.(int))
		}
	}
	i = &buffer
}

// Sum of all values in a given float slice
func (f *Float) Sum() interface{} {
	var sum float64
	for num := range f.Convert() {
		sum += num.(float64)
	}
	return sum
}

// Average of all values in a given float slice
func (f *Float) Average() interface{} {
	var sum float64
	for num := range f.Convert() {
		sum += num.(float64)
	}
	return sum
}

// Min of all values in a given float slice
func (f *Float) Min() interface{} {
	var min interface{} = <-f.Convert()
	for num := range f.Convert() {
		if num.(float64) < min.(float64) {
			min = num
		}
	}
	return min
}

// Max of all values in a given float slice
func (f *Float) Max() interface{} {
	var max interface{} = <-f.Convert()
	for num := range f.Convert() {
		if num.(float64) > max.(float64) {
			max = num
		}
	}
	return max
}

// Filter values from a given float slice
func (f *Float) Filter(v func(float64) bool) {
	var buffer Float = make(Float, 0, f.Len())
	for num := range f.Convert() {
		if v(num.(float64)) {
			buffer = append(buffer, num.(float64))
		}
	}
	f = &buffer
}

// Map the values of a Int slice to a given function
func (i *Integer) Map(f func(int) int) {
	var count int
	for num := range i.Convert() {
		(*i)[count] = f(num.(int))
		count++
	}
}

// Map the values of a float slice to a give function
func (f *Float) Map(v func(float64) float64) {
	var count int
	for num := range f.Convert() {
		(*f)[count] = v(num.(float64))
		count++
	}
}

// Distinct all values are distinct in a given stream of Ints
func (i *Integer) Distinct() interface{} {
	var mode map[int]bool = make(map[int]bool)
	for _, v := range *i {
		if _, ok := mode[v]; ok {
			return false
		}
		mode[v] = true
	}
	return true
}

// Peek the elements from the stream
func (i *Integer) Peek() interface{} {
	return (*i)[0]
}

// Sort the given stream
func (i *Integer) Sort() {
	sort.Ints(*i)
}

// ForEach does an operation on each element in the stream
func (i *Integer) ForEach(f func(int)) {
	for num := range i.Convert() {
		f(num.(int))
	}
}

// ToSlice convert a stream to an Integer slice
func (i *Integer) ToSlice() []int {
	return *i
}

// Skip the number of elements in the stream
func (i *Integer) Skip(num int) {
	*i = (*i)[num:]
}

// Distinct is true if all values in the stream are unique
func (f *Float) Distinct() interface{} {
	var mode map[float64]bool = make(map[float64]bool)
	for _, v := range *f {
		if _, ok := mode[v]; ok {
			return false
		}
		mode[v] = true
	}
	return true
}

//Peek the elements from the stream
func (f *Float) Peek() float64 {
	return (*f)[0]
}

// ForEach does an operation on each element in the stream
func (f *Float) ForEach(v func(float64)) {
	for num := range f.Convert() {
		v(num.(float64))
	}
}

// ToSlice converts a stream to its equivalent interface
func (f *Float) ToSlice() []float64 {
	return *f
}

// Skip the first n values in a given stream
func (f *Float) Skip(num int) {
	*f = (*f)[num:]
}
