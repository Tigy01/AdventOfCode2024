package main

import (
	"aoc/internal/day6"
	"fmt"
	"time"
)

func main() {
    start:= time.Now()
    fmt.Println("Advent OF Code 2024:\n")

    fmt.Println("Day6:")
    err:=day6.Run()
    if err != nil {
        fmt.Printf("err: %v\n", err)
    }
    fmt.Printf("Time Taken %v", time.Since(start).Microseconds())
}
