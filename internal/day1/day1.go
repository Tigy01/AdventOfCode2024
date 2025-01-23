package day1

import (
	"aoc/internal/library/filemanager"
	"fmt"
	"log"
	"slices"
	"strconv"
)

func Run() (output string) {
	//	lines, err := filemanager.ReadLines("./internal/day1/testInput.txt")
	lines, err := filemanager.ReadLines("./internal/day1/realInput.txt")
	if err != nil {
		log.Fatal("err: %v", err)
	}
	var left = []int{}
	var right = []int{}
	for _, line := range lines {
		var num_list, err = getNumbersFromLine(line)
		if err != nil {
			log.Fatal("err: %v", err)
		}
		left = append(left, num_list[0])
		right = append(right, num_list[1])
	}
	var pairs = getDistances(left, right)
	var frequencies = getFrequencies(left, right)
	var total_distances = 0
	for _, val := range pairs {
		total_distances += val
	}
	var total_frequencies = 0
	for _, val := range frequencies {
		total_frequencies += val
	}
	output += fmt.Sprintln("distances:", total_distances)
	output += fmt.Sprintln("frequencies:", total_frequencies)
	return
}

func getFrequencies(left_values []int, right_values []int) (frequencies []int) {
	var occurances = 0
	for _, left_num := range left_values {
		for i := len(right_values) - 1; i > -1; i-- {
            right_num :=right_values[i]
			if left_num == right_num {
				occurances += 1
			}
			right_values = slices.Delete(right_values, i, i+1)

		}
		frequencies = append(frequencies, left_num*occurances)
		occurances = 0
	}
	return frequencies
}

func getDistances(left_values []int, right_values []int) []int {
	slices.Sort(left_values)
	slices.Sort(right_values)
	var pairs = []int{}
	for i := range left_values {
		pairs = append(pairs, distance(right_values[i], left_values[i]))
	}
	return pairs
}

func getMin(list []int) int {
	minimum := list[0]
	for _, value := range list {
		if value < minimum {
			minimum = value
		}
	}
	return minimum
}

func distance(first int, second int) int {
	var big int
	if first >= second {
		big = first
	} else {
		big = second
	}
	var small int
	if first < second {
		small = first
	} else {
		small = second
	}
	return big - small
}

func getNumbersFromLine(line string) ([]int, error) {
	var currentNum = []rune{}
	var numlist = []int{}
	for i, char := range line {
		if IsNum(char) {
			currentNum = append(currentNum, char)
			if len(line)-1 != i {
				continue
			}

		}
		if string(currentNum) == "" {
			continue
		}

		var val, err = strconv.ParseInt(string(currentNum), 10, 64)
		if err != nil {
			return []int{}, err
		}
		numlist = append(numlist, int(val))
		var j = len(currentNum)
		for j > 0 {
			currentNum = slices.Delete(currentNum, j-1, j)
			j -= 1
		}
	}
	return numlist, nil
}

func IsNum(value rune) bool {
	if value < 48 || value > 57 {
		return false
	}
	return true
}
