package day11

import (
	"aoc/internal/day2"
	"aoc/internal/library/errors"
	"aoc/internal/library/filemanager"
	"fmt"
	"strconv"
)

//Inspired by an answer on reddit to use recursion and memoization.

func Run() (output string) {
	rawData := errors.Try(filemanager.ReadFullFile("./internal/day11/real.txt"))
	data := errors.Try(day2.GetNumbersFromLine(rawData))
	num := blinkNTimes(data, 75)

	output = fmt.Sprintf("Stones: %v\n", num)
	return
}

type MemoKey struct {
	stone, blinks int
}

var memos = map[MemoKey]int{}

func blinkNTimes(data []int, n int) (numStones int) {
	for _, v := range data {
		if stones, ok := memos[MemoKey{v, n}]; ok {
			numStones += stones
			return
		}
		numStones += blinkRecursive(v, n)
	}
	return
}

func blinkRecursive(value, n int) (numStones int) {
	if n == 0 {
		return 1
	}

	if stones, ok := memos[MemoKey{value, n}]; ok {
		numStones += stones
		return
	}

	if value == 0 {
		numStones = blinkRecursive(1, n-1)
	} else if stringValue := fmt.Sprint(value); len(stringValue)%2 == 0 {
		val1, val2 := splitStone(stringValue)
		numStones += blinkRecursive(val1, n-1)
		numStones += blinkRecursive(val2, n-1)
	} else {
		numStones = blinkRecursive(value*2024, n-1)
	}
	memos[MemoKey{value, n}] = numStones
	return
}

func splitStone(value string) (left, right int) {
	left = errors.Try(strconv.Atoi(value[:len(value)/2]))
	right = errors.Try(strconv.Atoi(value[len(value)/2:]))
	return
}
