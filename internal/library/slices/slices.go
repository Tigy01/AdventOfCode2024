package slices

func Swap[e any](data []e, first int, second int) []e {
	data[first], data[second] = data[second], data[first]
	return data 
}
