package filemanager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("An error occurred")
		fmt.Println(err)
		return nil, errors.New("error")
	}

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println("Failed to read file")
		fmt.Println(err)
		file.Close()
		return nil, errors.New("error")
	}

	file.Close()
	return lines, nil
}