package day2

import (
	"aoc/internal/day1"
	"aoc/internal/filemanager"
	"fmt"
	"slices"
	"strconv"
)

func Run() error {
	lines, err := filemanager.ReadLines("./internal/day2/realInput.txt")

	num_safe := 0
	num_safe_with_damp := 0
	if err != nil {
		return err
	}
	for _, line := range lines {
		var num_list, err = getNumbersFromLine(line)
		if err != nil {
			return err
		}

		if isSafe(num_list) {
			num_safe += 1
		}

		if isSafeWithDampen(num_list, false) {
			num_safe_with_damp += 1
		}
	}
	fmt.Println("Number Safe:", num_safe)
	fmt.Println("Number Safe With Dampening:", num_safe_with_damp)
	return nil
}

const INCREASING = 0
const DECREASING = 1
const UNKNOWN = 2

func isSafe(num_list []int) bool {
	if len(num_list) <= 1 {
		return false
	}

	var change = UNKNOWN
	for i, value := range num_list {
		if i == len(num_list)-1 {
			return true
		}

		var next_value = num_list[i+1]
		var delta = next_value - value
		var current_change = UNKNOWN
		if delta > 0 {
			current_change = INCREASING
		} else {
			current_change = DECREASING
		}

		var distance = delta
		if delta < 0 {
			distance *= -1
		}

		if distance < 1 || distance > 3 {
			return false
		}

		if change == UNKNOWN {
			change = current_change
		} else if current_change != change {
			return false
		}
	}
	return false
}

func isSafeWithDampen(num_list []int, dampened bool) bool {
	if len(num_list) <= 1 {
		return false
	}
	var change = UNKNOWN
	for i, value := range num_list {
		if i == len(num_list)-1 {
			return true
		}
		var next_value = num_list[i+1]
		var delta = next_value - value
		var current_change = UNKNOWN
		if delta > 0 {
			current_change = INCREASING
		} else {
			current_change = DECREASING
		}
		var distance = delta
		if delta < 0 {
			distance *= -1
		}

		if distance < 1 || distance > 3 {
			if dampened {
				return false
			}

			if dampen_list(num_list, i) {
				return true
			}

			return dampen_list(num_list, i+1)
		}
		//
		if change == UNKNOWN {
			change = current_change
		} else if current_change != change {
			if dampened {
				return false
			}

			if dampen_list(num_list, i) {
				return true
			}
			if dampen_list(num_list, i+1) {
				return true
			}
			if i == 1 {
				return dampen_list(num_list, i-1)
			}

			return false
		}

	}

	return false
}

func dampen_list(num_list []int, i int) bool {
	var dampened_list = []int{}

	dampened_list = append(dampened_list, num_list...)
	dampened_list = slices.Delete(dampened_list, i, i+1)
	var safe_with_dampen = isSafeWithDampen(dampened_list, true)
	return safe_with_dampen
}

func getNumbersFromLine(line string) (nums []int, err error) {
	var num_buff = []rune{}
	for i, char := range line {
		if day1.IsNum(char) {
			num_buff = append(num_buff, char)
			if len(line)-1 != i {
				continue
			}
		}

		if string(num_buff) == "" {
			continue
		}
		var val, err = strconv.ParseInt(string(num_buff), 10, 64)
		if err != nil {
			return []int{}, err
		}

		nums = append(nums, int(val))
		num_buff = []rune{}
	}

	return nums, nil

}
