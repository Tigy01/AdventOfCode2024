package day4

import (
	"aoc/internal/filemanager"
	"fmt"
)

type vec2 struct {
	x int
	y int
}

func Run() error {
	//	lines, err := filemanager.ReadLines("./internal/day4/testInput.txt")
	lines, err := filemanager.ReadLines("./internal/day4/realInput.txt")
	if err != nil {
		return err
	}

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
	fmt.Println("XMAS count:", matches)
	fmt.Println("X-MAS count:", crosses)
	return nil
}

func checkForX(lines []string, position vec2, token string) bool {
	if lines[position.y][position.x] != token[len(token)/2] {
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
	if position.y+direction.y < 0 || position.y+direction.y >= len(lines) {
		return ""
	}
	if position.x+direction.x < 0 || position.x+direction.x >= len(lines) {
		return ""
	}
	if position.y-direction.y < 0 || position.y-direction.y >= len(lines) {
		return ""
	}
	if position.x-direction.x < 0 || position.x-direction.x >= len(lines) {
		return ""
	}

	output += string(lines[position.y+direction.y][position.x+direction.x])
	output += string(lines[position.y][position.x])
	output += string(lines[position.y-direction.y][position.x-direction.x])

	return output
}

func checkDirections(lines []string, position vec2, token string) int {
	if lines[position.y][position.x] != token[len(token)-1] && lines[position.y][position.x] != token[0] {
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
		output += string(lines[position.y][position.x])
		if (position.y > 0 && offset.y < 0) || (position.y >= 0 && offset.y >= 0) {
			position.y += offset.y
		} else {
			break
		}
		if position.y == y_len {
			break
		}

		x_len = len(lines[position.y])
		if (position.x > 0 && offset.x < 0) || (position.x >= 0 && offset.x >= 0) {
			position.x += offset.x
		} else {
			break
		}
		if position.x == x_len {
			break
		}
	}
	return output
}
