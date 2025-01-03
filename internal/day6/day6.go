package day6

import (
	"aoc/internal/filemanager"
	"fmt"
)

type vec2 struct {
	x int
	y int
}

func (self vec2) eql(other vec2) bool {
	return self.x == other.x && self.y == other.y
}

var directions = map[string]vec2{
	"up":    {x: 0, y: -1},
	"down":  {x: 0, y: 1},
	"left":  {x: -1, y: 0},
	"right": {x: 1, y: 0},
}

func rotate90(dir vec2) vec2 {
	if dir.eql(directions["up"]) {
		return directions["right"]
	}
	if dir.eql(directions["down"]) {
		return directions["left"]
	}
	if dir.eql(directions["left"]) {
		return directions["up"]
	}
	if dir.eql(directions["right"]) {
		return directions["down"]
	}
	return vec2{}
}

func Run() error {
	lines, err := filemanager.ReadLines("./internal/day6/testInput.txt")
	if err != nil {
		return err
	}
	var new_lines = lines

	var direction = directions["up"]

	for true {
		location := getGuardLocation(new_lines)

		if (location.eql(vec2{x: -1, y: -1})) {
			break
		}

		update := offsetUntilObstical(location, direction, new_lines)

		new_lines = update

		direction = rotate90(direction)
	}

	var visited = 0
	for _, value := range new_lines {
		for _, char := range value {
			if char == 'X' {
				visited += 1
			}
		}
	}

    fmt.Println("visited:",visited)
	return nil
}

func offsetUntilObstical(location vec2, offset vec2, lines []string) []string {
	var line_changes = make([]string, len(lines))
	for i := range line_changes {
		line_changes[i] = lines[i]
	}
	var new_y = location.y
	var new_x = location.x
	for true {
		var line = line_changes[new_y]
		var updated_line = []rune(line)

		if line[new_x] == '#' {
			new_y -= offset.y
			new_x -= offset.x
			var old_line = line_changes[new_y]
			var updated_old_line = []rune(old_line)

			updated_old_line[new_x] = '^'
			line_changes[new_y] = string(updated_old_line)
			break
		}

		updated_line[new_x] = 'X'
		line_changes[new_y] = string(updated_line)

		if (new_y > 0 && offset.y < 0) || (new_y >= 0 && offset.y >= 0) {
			new_y += offset.y
		} else {
			break
		}

		if (new_x > 0 && offset.x < 0) || (new_x >= 0 && offset.x >= 0) {
			new_x += offset.x
		} else {
			break
		}

		if new_y == len(lines) {
			break
		}

		if new_x == len(line) {
			break
		}
	}
	return line_changes
}

func getGuardLocation(lines []string) vec2 {
	for y, line := range lines {
		for x, char := range line {
			if char == '^' {
				return vec2{x: x, y: y}
			}
		}
	}
	return vec2{-1, -1}
}
