package slices

import "fmt"

func Swap[e any](data []e, first int, second int) []e {
	data[first], data[second] = data[second], data[first]
	return data
}

func AppendAt[e any](data []e, index int, values ...e) ([]e, error) {
	if len(data) < int(index) {
		return data, fmt.Errorf("%v exceeds the maximum value for insertion %v", index, len(data))
	}
	if index < 0 {
		return data, fmt.Errorf("%v cannot be less than 0", index, len(data))
	}

	newSlice := make([]e, 0, len(data)+len(values))
	newSlice = append(newSlice, data[:index]...)
		newSlice = append(newSlice, values...)
	    newSlice = append(newSlice, data[index:]...)
//	for i := range index {
//		newSlice[i] = data[i]
//	}
//	for i, value := range values {
//		newSlice[index+i] = value
//	}
//	for i := index; i+len(values) < len(newSlice); i++ {
//		newSlice[i+len(values)] = data[i]
//	}
	return newSlice, nil
}
