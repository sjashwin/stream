package stream

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
// Map the values of a float slice to a given function
func (i *Integer) Map(f func(int) int) {
	var count int
	for num := range i.Convert() {
		i.Values[count] = f(num)
		count++
	}
>>>>>>> Stream Functionality: Adds Support For Maps
}
