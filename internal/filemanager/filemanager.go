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

func ReadFullFile(path string) (output string, err error) {
	data, err := os.Open(path)
	if err != nil {
		return "", err
	}
	buff := make([]byte, 1024, 1024)
	bytes, err := data.Read(buff)
	for err == nil {
		output += string(buff[0:bytes])
		bytes, err = data.Read(buff)
	}
    return output, nil
}
