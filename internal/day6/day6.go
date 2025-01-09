package day6

import (
	"aoc/internal/library/filemanager"
	"aoc/internal/library/vectors"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
)

type vec2 = vectors.Vec2

var directions = map[string]vec2{
	"up":    {X: 0, Y: -1},
	"down":  {X: 0, Y: 1},
	"left":  {X: -1, Y: 0},
	"right": {X: 1, Y: 0},
}

func rotate90(dir vec2) vec2 {
	if dir == directions["up"] {
		return directions["right"]
	}
	if dir == directions["down"] {
		return directions["left"]
	}
	if dir == directions["left"] {
		return directions["up"]
	}
	if dir == directions["right"] {
		return directions["down"]
	}
	return vec2{}
}

func Run() (output string) {

	//	err := runOnDataSet("./internal/day6/generator/test.txt")
	output, err := runOnFile("./internal/day6/real.txt")
	if err != nil {
		log.Fatal("err: %v\n", err)
	}
	return output
}

func runOnFile(fileName string) (string, error) {
	lines, err := filemanager.ReadLines(fileName)
	if err != nil {
		return "", err
	}
	//	var newLines = lines
	var newLines = convertLines(lines)

	var direction = directions["up"]

	location := getGuardLocation(newLines)
	if (location == vec2{X: -1, Y: -1}) {
		return "", nil
	}

	loopingPositions := checkForAvaliableLoop(location, direction, clone(newLines))
	for true {
		newLines, location = offsetUntilObstical(location, direction, newLines)
		direction = rotate90(direction)

		if (location == vec2{X: -1, Y: -1}) {
			break
		}

	}

	var visited = 0
	for _, value := range newLines {
		for _, v := range value {
			if v == 1 {
				visited += 1
			}
		}
	}

	if visited == 0 { //|| avaliableObstructions == 0 {
		return "", badUpdate
	}
	output := fmt.Sprintln("visited:", visited)
	output += fmt.Sprintln("obstructions:", loopingPositions)
	return output, nil
}

const OBSTICAL = 3
const GUARD = 2
const VISITED = 1

func convertLines(lines []string) [][]int {
	convertedLines := [][]int{}
	for y, line := range lines {
		convertedLines = append(convertedLines, []int{})
		for _, char := range line {
			value := 0
			switch char {
			case '#':
				value = OBSTICAL
				break
			case '^':
				value = GUARD
				break
			default:
				value = 0
				break
			}
			convertedLines[y] = append(convertedLines[y], value)
		}
	}
	return convertedLines
}

func clone[e any](x [][]e) [][]e {
	y := [][]e{}
	for _, v := range x {
		y = append(y, slices.Clone(v))
	}
	return y
}

var badUpdate = errors.New("badUpdate")

func runOnDataSet(fileName string) (output string, err error) {
	i := 0
	path := fileName

	for i < 100 {
		newOutput, err := runOnFile(path)
		if err != nil {
			if errors.Is(err, badUpdate) {
				os.Remove(path)
			}
		} else {
			output += fmt.Sprintln(newOutput, path, "\n")
		}

		i += 1
		path = fmt.Sprintf("./internal/day6/generator/test%v.txt", i)
	}
	return output, nil
}

type pair struct {
	location, direction vec2
}

func checkForAvaliableLoop(location, direction vec2, lines [][]int) (loopingPositions int) {
	loopingPlacements := []vec2{}
	for {
		if (location == vec2{-1, -1}) {
			break
		}
		nextLines, projectedLocation := offsetUntilObstical(location, direction, clone(lines))

		obstacleLocation := location

		for obstacleLocation != projectedLocation { //checks a line between start pos and end position
			var visitedPairs = []pair{}
			obstacleLocation = obstacleLocation.Add(direction)

			if !obstacleLocation.IsInRange(vec2{0, 0}, vec2{len(lines[0]), len(lines)}) {
				break
			}

			linesWithObstical := clone(lines) //lines after placing obstical
			linesWithObstical[obstacleLocation.Y][obstacleLocation.X] = 3

			locationWithObstical := location //location of guard prior to placement of a new ob
			newDir := direction              //spin without affecting the original placement

			numSpinsWithoutMovement := 0 //prevents infinite spinning
			moves := 0
			for (locationWithObstical != vec2{-1, -1}) { //checks all positions after obstical is placed until hits the void
				if slices.Contains(visitedPairs, pair{locationWithObstical, newDir}) {
					if !slices.Contains(loopingPlacements, obstacleLocation) {
						loopingPlacements = append(loopingPlacements, obstacleLocation)
						//						fmt.Println("loop at:", obstacleLocation)
					}
					break
				}

				previousLocal := locationWithObstical

				linesWithObstical, locationWithObstical = offsetUntilObstical(
					locationWithObstical, newDir, linesWithObstical,
				) //project foreward one step

				if locationWithObstical != previousLocal { //if the guard isnt boxed in
					visitedPairs = append(visitedPairs, pair{previousLocal, newDir})
				} else { //check if he will be stuck infinitely
					numSpinsWithoutMovement += 1
					if numSpinsWithoutMovement > 6 {
						break
					}
				}

				moves += 1

				newDir = rotate90(newDir)
			}
		}
		lines = nextLines
		direction = rotate90(direction)
		location = projectedLocation
	}
	return len(loopingPlacements)
}

func offsetUntilObstical(location vec2, offset vec2, lines [][]int) (newLines [][]int, newPosition vec2) {
	newPosition = location
	newLines = lines
	for true {
		var line = newLines[newPosition.Y]

		if line[newPosition.X] == OBSTICAL {
			newPosition = newPosition.Sub(offset)
			newLines[newPosition.Y][newPosition.X] = GUARD
			return newLines, newPosition
		}

		newLines[newPosition.Y][newPosition.X] = VISITED
		newPosition = newPosition.Add(offset)
		if !newPosition.IsInRange(vec2{0, 0}, vec2{len(line), len(newLines)}) {
			return newLines, vec2{-1, -1}
		}
	}
	return
}

func getGuardLocation(lines [][]int) vec2 {
	for y, line := range lines {
		if x := slices.Index(line, GUARD); x >= 0 {
			return vec2{x, y}
		}
	}
	return vec2{-1, -1}
}
