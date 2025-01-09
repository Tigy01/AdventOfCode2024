package main

import (
	"aoc/internal/day1"
	"aoc/internal/day2"
	"aoc/internal/day3"
	"aoc/internal/day4"
	"aoc/internal/day5"
	"aoc/internal/day6"
	"aoc/internal/day7"
	"aoc/internal/day8"
	"aoc/internal/day9"

	"aoc/internal/library/errors"
	"os"
	"runtime/pprof"

	"fmt"
	"time"
)

func main() {
	file := errors.Try(os.Create("current.prof"))

	pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()

	fmt.Println("Advent Of Code 2024:\n")
	dayCount := 9
	doneChan := make(chan bool, dayCount)
	go runDay(1, day1.Run, doneChan)
	go runDay(2, day2.Run, doneChan)
	go runDay(3, day3.Run, doneChan)
	go runDay(4, day4.Run, doneChan)
	go runDay(5, day5.Run, doneChan)
	go runDay(6, day6.Run, doneChan)
	go runDay(7, day7.Run, doneChan)
	go runDay(8, day8.Run, doneChan)
	go runDay(9, day9.Run, doneChan)
	for range dayCount {
		<-doneChan
	}
}

func runDay(dayNum int, run func() string, done chan bool) {
	start := time.Now()
	output := run()
    final:= fmt.Sprintf("Day %v:\n", dayNum)
	final+=fmt.Sprintf(output)
	final+=fmt.Sprintf("Time Taken: %v\n\n", time.Since(start).Milliseconds())
    fmt.Println(final)
	done <- true
}
