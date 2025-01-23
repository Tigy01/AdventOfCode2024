package day4

import (
	"aoc/internal/library/errors"
	"aoc/internal/library/filemanager"
	"aoc/internal/library/vectors"
	"fmt"
)

type vec2 = vectors.Vec2

func Run() (output string) {
	//	lines, err := filemanager.ReadLines("./internal/day4/testInput.txt")
	lines := errors.Try(filemanager.ReadLines("./internal/day4/realInput.txt"))

	var matches = 0
	var crosses = 0
	for y, line := range lines {
		for x := range line {
			matches += checkDirections(lines, vec2{x, y}, "XMAS")
			if checkForX(lines, vec2{x, y}, "MAS") {
				crosses += 1
			}
		}
	}
	output += fmt.Sprintln("XMAS count:", matches)
	output += fmt.Sprintln("X-MAS count:", crosses)
	return
}

func checkForX(lines []string, position vec2, token string) bool {
	if lines[position.Y][position.X] != token[len(token)/2] {
		return false
	}

	var numFound = 0
	var south_east = getOffsetFromCenter(lines, position, vec2{1, 1})
	var south_west = getOffsetFromCenter(lines, position, vec2{-1, 1})
	if south_east == token || isReversed(south_east, token) {
		numFound += 1
	}
	if south_west == token || isReversed(south_west, token) {
		numFound += 1
	}
	return numFound == 2
}

func isReversed(str, token string) bool {
	var i = len(token) - 1
	var j = 0
	if len(str) != len(token) {
		return false
	}
	for true {
		if str[j] != token[i] {
			return false
		}
		if i == 0 {
			break
		}
		i -= 1
		j += 1
	}
	return true
}

func getOffsetFromCenter(lines []string, position, direction vec2) string {
	output := ""
	if position.Y+direction.Y < 0 || position.Y+direction.Y >= len(lines) {
		return ""
	}
	if position.X+direction.X < 0 || position.X+direction.X >= len(lines) {
		return ""
	}
	if position.Y-direction.Y < 0 || position.Y-direction.Y >= len(lines) {
		return ""
	}
	if position.X-direction.X < 0 || position.X-direction.X >= len(lines) {
		return ""
	}

	output += string(lines[position.Y+direction.Y][position.X+direction.X])
	output += string(lines[position.Y][position.X])
	output += string(lines[position.Y-direction.Y][position.X-direction.X])

	return output
}

func checkDirections(lines []string, position vec2, token string) int {
	if lines[position.Y][position.X] != token[len(token)-1] && lines[position.Y][position.X] != token[0] {
		return 0
	}

	var numFound = 0
	var directions = [4]string{}
	directions[0] = getOffset(lines, 4, position, vec2{1, 1})  //South East
	directions[1] = getOffset(lines, 4, position, vec2{1, -1}) //North West
	directions[2] = getOffset(lines, 4, position, vec2{1, 0})  //East
	directions[3] = getOffset(lines, 4, position, vec2{0, 1})  //South

	for _, dir := range directions {
		if dir == token || isReversed(dir, token) {
			numFound += 1
		}
	}
	return numFound
}

func getOffset(lines []string, length int, position, offset vec2) (output string) {
	var y_len = len(lines)
	var x_len = 0

	for range length {
		output += string(lines[position.Y][position.X])
		if (position.Y > 0 && offset.Y < 0) || (position.Y >= 0 && offset.Y >= 0) {
			position.Y += offset.Y
		} else {
			break
		}
		if position.Y == y_len {
			break
		}

		x_len = len(lines[position.Y])
		if (position.X > 0 && offset.X < 0) || (position.X >= 0 && offset.X >= 0) {
			position.X += offset.X
		} else {
			break
		}
		if position.X == x_len {
			break
		}
	}
	return output
}
