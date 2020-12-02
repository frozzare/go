package file

import (
	"bufio"
	"os"
)

// ReadLines reads a file and return the lines or a error.
func ReadLines(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var list []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return list, nil
}
