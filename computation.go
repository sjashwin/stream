package stream

import "sort"

// Sum all the values in a given stream
func (i *Integer) Sum() int {
	var sum int
	for num := range i.Convert() {
		sum += num
	}
	return sum
}

// Average of all values in a given stream
func (i *Integer) Average() int {
	var sum int
	for num := range i.Convert() {
		sum += num
	}
	return sum / len(i.Values)
}

// Min of all values in a given stream
func (i *Integer) Min() int {
	var min int = <-i.Convert()
	for num := range i.Convert() {
		if num < min {
			min = num
		}
	}
	return min
}

//Max of all the values in a given stream
func (i *Integer) Max() int {
	var max int = <-i.Convert()
	for num := range i.Convert() {
		if num > max {
			max = num
		}
	}
	return max
}

// Filter values from a given integer slice
func (i *Integer) Filter(f func(int) bool) {
	var buffer []int = make([]int, 0, len(i.Values))
	for num := range i.Convert() {
		if f(num) {
			buffer = append(buffer, num)
		}
	}
	i.Values = buffer
}

// Sum of all values in a given float slice
func (f *Float) Sum() float64 {
	var sum float64
	for num := range f.Convert() {
		sum += num
	}
	return sum
}

// Average of all values in a given float slice
func (f *Float) Average() float64 {
	var sum float64
	for num := range f.Convert() {
		sum += num
	}
	return sum
}

// Min of all values in a given float slice
func (f *Float) Min() float64 {
	var min float64 = <-f.Convert()
	for num := range f.Convert() {
		if num < min {
			min = num
		}
	}
	return min
}

// Max of all values in a given float slice
func (f *Float) Max() float64 {
	var max float64 = <-f.Convert()
	for num := range f.Convert() {
		if num > max {
			max = num
		}
	}
	return max
}

// Filter values from a given float slice
func (f *Float) Filter(v func(float64) bool) {
	var buffer []float64 = make([]float64, 0, len(f.Values))
	for num := range f.Convert() {
		if v(num) {
			buffer = append(buffer, num)
		}
	}
	f.Values = buffer
}

// Map the values of a Int slice to a given function
func (i *Integer) Map(f func(int) int) {
	var count int
	for num := range i.Convert() {
		i.Values[count] = f(num)
		count++
	}
}

// Map the values of a float slice to a give function
func (f Float) Map(v func(float64) float64) {
	var count int
	for num := range f.Convert() {
		f.Values[count] = v(num)
		count++
	}
}

// Distinct all values are distinct in a given stream of Ints
func (i *Integer) Distinct() bool {
	var mode map[int]bool = make(map[int]bool)
	for _, v := range i.Values {
		if _, ok := mode[v]; ok {
			return false
		}
		mode[v] = true
	}
	return true
}

// Peek the elements from the stream
func (i *Integer) Peek() int {
	return i.Values[0]
}

// Sort the given stream
func (i *Integer) Sort() {
	sort.Ints(i.Values)
}

// ForEach does an operation on each element in the stream
func (i *Integer) ForEach(f func(int)) {
	for num := range i.Convert() {
		f(num)
	}
}

// ToSlice convert a stream to an Integer slice
func (i *Integer) ToSlice() []int {
	return i.Values
}

// Skip the number of elements in the stream
func (i *Integer) Skip(num int) {
	i.Values = i.Values[num:]
}

// Distinct is true if all values in the stream are unique
func (f *Float) Distinct() bool {
	var mode map[float64]bool = make(map[float64]bool)
	for _, v := range f.Values {
		if _, ok := mode[v]; ok {
			return false
		}
		mode[v] = true
	}
	return true
}

//Peek the elements from the stream
func (f *Float) Peek() float64 {
	return f.Values[0]
}

// ForEach does an operation on each element in the stream
func (f *Float) ForEach(v func(float64)) {
	for num := range f.Convert() {
		v(num)
	}
}

// ToSlice converts a stream to its equivalent interface
func (f Float) ToSlice() []float64 {
	return f.Values
}

// Skip the first n values in a given stream
func (f *Float) Skip(num int) {
	f.Values = f.Values[num:]
}
