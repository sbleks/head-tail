package headtail

import (
	"bufio"
	"os"
)

// Head returns the top n lines of a file like the UNIX command.
func Head(path string, n int) ([]string, error) {
	lines := make([]string, 0, n)

	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(file)

	for i := 0; fileScanner.Scan() && i < n; i++ {
		lines = append(lines, fileScanner.Text())
	}

	return lines, nil
}

// Tail returns the bottom n lines of a file like the UNIX command.
func Tail(path string, n int) ([]string, error) {
	var lines []string

	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	if n > len(lines) {
		return lines, nil
	}

	return lines[len(lines)-n:], nil
}
