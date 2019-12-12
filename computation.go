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

<<<<<<< HEAD
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
=======
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
>>>>>>> Stream Functionalities: Adds More Support To Integer
}
