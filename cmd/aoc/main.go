package main

import (
	"aoc/internal/day1"
	"aoc/internal/day2"
	"aoc/internal/day3"
	"aoc/internal/day4"
	"aoc/internal/day5"
	"aoc/internal/day6"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Advent Of Code 2024:\n")
	runDay(1, day1.Run)
	runDay(2, day2.Run)
	runDay(3, day3.Run)
	runDay(4, day4.Run)
	runDay(5, day5.Run)
	runDay(6, day6.Run)
}

func runDay(dayNum int, run func() error) {
	start := time.Now()
	fmt.Printf("Day %v:\n", dayNum)
	err := run()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("Time Taken: %v\n\n", time.Since(start).Microseconds())
}
