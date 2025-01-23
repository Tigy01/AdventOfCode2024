package day10

import (
	"aoc/internal/library/errors"
	"aoc/internal/library/filemanager"
	"aoc/internal/library/vectors"
	"fmt"
	"slices"
	"strconv"
)

type vec2 = vectors.Vec2

var directions = map[string]vec2{
	"up":    {X: 0, Y: -1},
	"down":  {X: 0, Y: 1},
	"left":  {X: -1, Y: 0},
	"right": {X: 1, Y: 0},
}

func Run() (output string) {
	lines := errors.Try(filemanager.ReadLines("./internal/day10/real.txt"))
	data := convertToReadableData(lines)
	trailHeads := getTrailHeads(data)
	nines := 0
	trails := 0
	for _, head := range trailHeads {
		nines += getNumberOfNines(data, evaluateAllPositions(data, head, []vec2{}, false))
		trails += getNumberOfNines(data, evaluateAllPositions(data, head, []vec2{}, true))
	}
	output += fmt.Sprintf("Number of Nines: %v\n", nines)
	output += fmt.Sprintf("Number of Trails: %v\n", trails)
	return
}

func getNumberOfNines(data [][]int, positions []vec2) (nines int) {
	for _, position := range positions {
		if data[position.Y][position.X] == 9 {
			nines += 1
		}
	}
	return
}

func convertToReadableData(lines []string) (data [][]int) {
	data = make([][]int, len(lines))
	for y, line := range lines {
		dataLine := make([]int, len(line))
		for x, char := range line {
			if val, err := strconv.ParseInt(string(char), 10, 64); err != nil {
				dataLine[x] = -2
			} else {
				dataLine[x] = int(val)
			}
		}
		data[y] = dataLine
	}
	return
}

func getTrailHeads(data [][]int) (trailheads []vec2) {
	for y, line := range data {
		for x, val := range line {
			if val == 0 {
				trailheads = append(trailheads, vec2{x, y})
			}
		}
	}
	return
}

func evaluateAllPositions(data [][]int, start vec2, previousPositions []vec2, pt2 bool) []vec2 {
	if data[start.Y][start.X] == 9 {
		if slices.Contains(previousPositions, start) {
			return previousPositions
		} else {
			previousPositions = append(previousPositions, start)
			return previousPositions
		}
	}
	nextPositions := getNextPositions(data, start)
	for _, pos := range nextPositions {
		if slices.Contains(previousPositions, pos) && !pt2 {
			continue
		}
		previousPositions = append(previousPositions, pos)
		previousPositions = evaluateAllPositions(data, pos, previousPositions, pt2)
	}
	return previousPositions
}

func getNextPositions(data [][]int, position vec2) []vec2 {
	validPositions := make([]vec2, 0, 4)
	currVal := data[position.Y][position.X]
	for _, dir := range directions {
		nextPos := position.Add(dir)
		if !nextPos.IsInRange(vec2{0, 0}, vec2{len(data[position.Y]), len(data)}) {
			continue
		}
		nextVal := data[nextPos.Y][nextPos.X]
		if nextVal-currVal > 0 && nextVal-currVal < 2 {
			validPositions = append(validPositions, nextPos)
		}
	}
	return validPositions
}
