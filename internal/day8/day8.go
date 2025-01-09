package day8

import (
	"aoc/internal/library/filemanager"
	mystrings "aoc/internal/library/strings"
	"aoc/internal/library/vectors"
	"fmt"
	"log"
	"strings"
)

type vec2 = vectors.Vec2

type coordinateRange struct {
	lower vec2
	upper vec2
}

func Run() (output string) {

	lines, err := filemanager.ReadLines("./internal/day8/real.txt")
	if err != nil {
		log.Fatal(err)
	}

	stations := readStationMap(lines)
	antinodes := calculateAntiNodeLocations(stations, coordinateRange{}, false)
	antinodesP2 := calculateAntiNodeLocations(
		stations,
		coordinateRange{
			vec2{X: 0, Y: 0},
			vec2{X: len(lines[0]), Y: len(lines)}},
		true,
	)
	lines = placeAntiNodes(antinodes, lines)
	antinodecount := 0
	for _, line := range lines {
		antinodecount += strings.Count(line, "#")
	}
	lines = placeAntiNodes(antinodesP2, lines)
	antinodecountP2 := 0
	for _, line := range lines {
		for _, v := range line {
			if mystrings.IsAlphaNumeric(v) || v == '#' {
				antinodecountP2 += 1
			}
		}
	}
	output += fmt.Sprintf("Anti nodes: %v\n", antinodecount)
	output += fmt.Sprintf("Anti nodes P2: %v\n", antinodecountP2)
	return
}

func placeAntiNodes(antinodes map[rune][]vec2, lines []string) []string {
	for key, locations := range antinodes {
		for _, loc := range locations {
			if loc.Y >= len(lines) || loc.Y < 0 {
				continue
			}
			if loc.X >= len(lines[loc.Y]) || loc.X < 0 {
				continue
			}
			if lines[loc.Y][loc.X] == byte(key) {
				continue
			}
			lines = mystrings.PlaceCharAtCoordinateInPlace('#', loc, lines)
		}
	}
	return lines
}

func readStationMap(lines []string) (stations map[rune][]vec2) {
	stations = map[rune][]vec2{}
	for y, line := range lines {
		for x, char := range line {
			if mystrings.IsAlphaNumeric(char) {
				list, ok := stations[char]
				if !ok {
					list = []vec2{}
				}
				list = append(list, vec2{X: x, Y: y})
				stations[char] = list
			}
		}
	}
	return
}

func calculateAntiNodeLocations(stations map[rune][]vec2, coordRange coordinateRange, pt2 bool) (antinodes map[rune][]vec2) {
	antinodes = map[rune][]vec2{}
	for key, locations := range stations {
		for i := 0; i < len(locations)-1; i++ {
			location := locations[i]
			for _, next_location := range locations[i+1:] {
				offset := next_location.Sub(location)
				antinodeLocations, ok := antinodes[key]
				if !ok {
					antinodeLocations = []vec2{}
				}
				beforeAntiNode := location.Sub(offset)
				afterAntiNode := next_location.Add(offset)
				if pt2 {
					for beforeAntiNode.IsInRange(coordRange.lower, coordRange.upper) {
						antinodeLocations = append(antinodeLocations, beforeAntiNode)
						beforeAntiNode = beforeAntiNode.Sub(offset)
					}
					for afterAntiNode.IsInRange(coordRange.lower, coordRange.upper) {
						antinodeLocations = append(antinodeLocations, afterAntiNode)
						afterAntiNode = afterAntiNode.Add(offset)
					}
				} else {
					antinodeLocations = append(antinodeLocations, beforeAntiNode)
					antinodeLocations = append(antinodeLocations, afterAntiNode)
				}
				antinodes[key] = antinodeLocations
			}
		}
	}
	return
}
