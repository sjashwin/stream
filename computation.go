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
