package strings

import (
	"aoc/internal/library/vectors"
	"slices"
)

type vec2 = vectors.Vec2

func PlaceCharAtCoordinate(char rune, location vec2, lines []string) []string {
	lines = slices.Clone(lines)
	newLine := lines[location.Y]
	newLine = newLine[:location.X] + string(char) + newLine[location.X+1:]
	lines[location.Y] = newLine
	return lines
}

func PlaceCharAtCoordinateInPlace(char rune, location vec2, lines []string) []string {
	newLine := lines[location.Y]
	newLine = newLine[:location.X] + string(char) + newLine[location.X+1:]
	lines[location.Y] = newLine
	return lines
}

func IsAlphaNumeric(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}
