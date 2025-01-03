package filemanager

import (
	"bufio"
	"os"
)


func ReadLines(path string) ([]string, error) {
	lines := make([]string, 0)
	data, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}
	fileScanner := bufio.NewScanner(data)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}
	return lines, nil
}


