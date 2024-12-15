package filemanager

import (
	"bufio"
	"encoding/json"
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

func WriteJSON(path string, data interface{}/*any*/) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if (err != nil) {
		file.Close()
		return errors.New("failed to convert data to json")
	}

	file.Close()
	return nil
}