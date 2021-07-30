package Integer

// Filter values from a given integer slice
func (i *Integer) Filter(f func(int) bool) *Integer {
	var buffer = make([]int, 0, len(i.values))
	for num := range i.values {
		if f(num) {
			buffer = append(buffer, num)
		}
	}
	i.values = buffer

	return i
}
