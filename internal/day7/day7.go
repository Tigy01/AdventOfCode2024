package day7

import (
	"aoc/internal/day2"
	"aoc/internal/library/filemanager"
	"fmt"
	"log"
	"slices"
)

func Run() (output string) {
	lines, err := filemanager.ReadLines("./internal/day7/real.txt")
	if err != nil {
		log.Fatal("err: %v\n", err)
	}
	runProcess := func(useConcat bool) int {
		total := 0
		resultChan := make(chan int, len(lines))
		for _, line := range lines {
			go processLine(line, useConcat, resultChan)
		}
		received := 0
		for received != len(lines) {
			total += <-resultChan
			received += 1
		}
		return total
	}

	total := runProcess(false)
	output += fmt.Sprint("part 1:", total, "\n")
	total = runProcess(true)
	output += fmt.Sprint("Part 2:", total, "\n")
	return
}

func processLine(line string, useConcat bool, resultChan chan int) {
	values, err := day2.GetNumbersFromLine(line)
	total := values[0]
	values = values[1:]
	if err != nil {
		log.Fatal("error: %v", err)
	}
	if good := evaluateBranches(values[0], values[1], total, values[2:], useConcat); good {
		resultChan <- total
		return
	}
	resultChan <- 0
}

func concatNum(first, second int) int {
	secondCpy := second
	factor := 1
	for secondCpy > 0 {
		secondCpy /= 10
		factor *= 10
	}
	result := first*factor + second
	return result
}

func evaluateBranches(firstNum, secondNum, total int, remaining []int, useConcat bool) (matchFound bool) {
	keys := [3]rune{'+', '*', '|'}
	vals := [3]int{0, 0, 0}
	for i, op := range keys {
		switch op {
		case '+':
			vals[i] = firstNum + secondNum
			break
		case '*':
			vals[i] = firstNum * secondNum
			break
		}
		if useConcat {
			concated := concatNum(firstNum, secondNum)
			vals[2] = concated
		}
	}

	if len(remaining) != 0 {
		branchResults := []bool{}
		for _, result := range vals {
			branchResult := evaluateBranches(result, remaining[0], total, remaining[1:], useConcat)
			branchResults = append(branchResults, branchResult)
		}
		return slices.Contains(branchResults, true)
	}

	for _, result := range vals {
		if result == total {
			matchFound = true
		}
	}

	return
}
