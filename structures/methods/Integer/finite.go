package Integer

import "math"

// Sum all the values in a given stream
func (i *Integer) Sum() int {
	var sum int
	for num := range i.values {
		sum += num
	}
	return sum
}

// Average of all values in a given stream
func (i *Integer) Average() int {
	var sum int
	for num := range i.values {
		sum += num
	}
	return sum / len(i.values)
}

// Min of all values in a given stream
func (i *Integer) Min() int {
	var min = math.MinInt64
	for num := range i.values {
		if num < min {
			min = num
		}
	}
	return min
}

//Max of all the values in a given stream
func (i *Integer) Max() int {
	var max = math.MaxInt64
	for num := range i.values {
		if num > max {
			max = num
		}
	}
	return max
}
