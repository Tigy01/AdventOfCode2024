package main

import (
	"aoc/internal/day1"
	"aoc/internal/day2"
	"aoc/internal/day3"
	"aoc/internal/day6"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Advent OF Code 2024:\n")

	start := time.Now()
	fmt.Println("Day1:")
	err := day1.Run()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("Time Taken: %v\n\n", time.Since(start).Microseconds())

	start = time.Now()
	fmt.Println("Day2:")
	err = day2.Run()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("Time Taken: %v\n\n", time.Since(start).Microseconds())

	start = time.Now()
	fmt.Println("Day3:")
	err = day3.Run()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("Time Taken: %v\n\n", time.Since(start).Microseconds())

	start = time.Now()
	fmt.Println("Day6:")
	err = day6.Run()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("Time Taken: %v\n\n", time.Since(start).Microseconds())
}
