package day9

import (
	"aoc/internal/library/errors"
	"aoc/internal/library/filemanager"
	"aoc/internal/library/slices"
	"aoc/internal/library/vectors"
	"fmt"
	goSlices "slices"
	"strconv"
)

type vec2 = vectors.Vec2

func Run() (output string) {
	data := errors.Try(filemanager.ReadFullFile("./internal/day9/real.txt"))
	chunks, spaces := processData(data)
	unalignedLayout := createLayout(chunks, spaces)
	alignedLayout := alignLayout(goSlices.Clone(unalignedLayout))
	chunkAlignedLayout := alignLayoutGrouped(chunks, unalignedLayout)
	checksum := calculateChecksum(alignedLayout)
	chunked_checksum := calculateChecksum(chunkAlignedLayout)
	output = fmt.Sprintf("Aligned Mem Checksum: %v \n", checksum)
	output += fmt.Sprintf("Chunked Mem Checksum: %v \n", chunked_checksum)
	return
}

func calculateChecksum(alignedLayout []int) (checksum int) {
	for i, value := range alignedLayout {
		if value == -1 {
			continue
		}
		checksum += i * value
	}
	return
}

func createLayout(chunks []int, spaces []int) (layout []int) {
	layout = make([]int, 0, 2048)
	for i, size := range chunks {
		for range size {
			layout = append(layout, i)
		}
		if i < len(spaces) {
			for range spaces[i] {
				layout = append(layout, -1)
			}
		}
	}
	return layout
}

func alignLayout(rawLayout []int) []int {
	lastKnownNum := len(rawLayout) - 1
	for i, val := range rawLayout {
		if val == -1 {
			for j := lastKnownNum; j > i; j-- {
				if rawLayout[j] > -1 {
					lastKnownNum = j
					slices.Swap(rawLayout, i, j)
					break
				}
			}
		}
	}
	return rawLayout
}

func alignLayoutGrouped(chunks, rawLayout []int) []int {
	data := vec2{-1, -1}
	for dataVal := len(chunks) - 1; dataVal > -1; dataVal-- {
		data.X = goSlices.Index(rawLayout, dataVal)
		data.Y = lastIndex(rawLayout[data.X:], dataVal) + data.X //speeds up by narrowing search range

		spaces := findFirstEmptyChunk(rawLayout, lenthOfChunk(data))
		if data.X > spaces.X {
			replaceRange(rawLayout, dataVal, spaces.X, spaces.Y+1)
			replaceRange(rawLayout, -1, data.X, data.Y+1)
		}
	}
	return rawLayout
}

func findFirstEmptyChunk(data []int, length int) vec2 {
	coords := vec2{-1, -1}
	for i := 0; i < len(data); i++ {
		if data[i] == -1 {
			if coords.X == -1 {
				coords.X = i
				coords.Y = i
			} else {
				coords.Y = i
			}
			if lenthOfChunk(coords) == length {
				break
			}
			continue
		}
		if coords.X != -1 {
			coords = vec2{-1, -1}
		}
	}
	return coords
}

func replaceRange[e any](slice []e, val e, i, j int) {
	for i < j {
		slice[i] = val
		i++
	}
}

func lenthOfChunk(vec vec2) int {
	return vec.Y + 1 - vec.X
}

func lastIndex[e comparable](s []e, value e) int {
	for i := len(s) - 1; i > -1; i-- {
		if s[i] == value {
			return i
		}
	}
	return -1
}

func processData(data string) (chuncks []int, spaces []int) {
	chuncks, spaces = make([]int, 0, 512), make([]int, 0, 512)
	for i, c := range data {
		size, err := strconv.ParseInt(string(c), 10, 64)
		if err != nil {
			return
		}
		if i%2 == 0 {
			chuncks = append(chuncks, int(size))
		} else {
			spaces = append(spaces, int(size))
		}
	}
	return
}
