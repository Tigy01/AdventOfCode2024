package main

import (
	"fmt"
	"math/rand"
	"os"
	"slices"
)

func main() {
    for i := 0; i < 100; i++ {
        file, err := getNewFile()
        if err != nil {
            fmt.Printf("err: %v\n", err)
            return
        }
        data := getRandData(10, 10, .90)
        data = placePlayer(data, 10, 10)
        for _, v := range data {
            fmt.Println(v)
        }
        if err = fillFile(file, data); err != nil {
            fmt.Printf("err: %v\n", err)
            return
        }
    }
}

func placePlayer(data []string, x_max, y_max int) []string {
	x := (x_max / 5) + rand.Intn(x_max/2)
	y := (y_max / 5) + rand.Intn(y_max/2)

	data = slices.Clone(data)
	newLine := []rune(data[y])
	newLine[x] = '^'
	data[y] = string(newLine)
	return data
}

func fillFile(file *os.File, data []string) error {
    for _, line := range data {
        _, err := file.WriteString(line) 
        if err != nil {
            return err
        }
        _, err = file.WriteString("\n") 
        if err != nil {
            return err
        }
    }
	return nil
}

func getRandData(x_max, y_max int, chance float32) []string {
	lines := make([]string, y_max)
	for y := range y_max {
		line := ""
		for i := 0; i < x_max; i++ {
			random := rand.Float32()
			if random > chance {
				line += "#"
				continue
			}
			line += "."
		}
		lines[y] = line
	}
	return lines
}

func getNewFile() (*os.File, error) {
	i := 0
	fileName := fmt.Sprintf("./test%v.txt", "")
	for {
		file, err := os.Open(fileName)

		if err == nil {
			i += 1
			fileName = fmt.Sprintf("./test%v.txt", i)

			if clerr := file.Close(); clerr != nil {
				return nil, clerr
			}
			continue
		}
		break
	}
	f, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	return f, nil
}
