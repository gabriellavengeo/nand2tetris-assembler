package utils

import (
	"bufio"
	"os"
	"strings"
)

// ReadLines reads and returns all the lines from a file
func ReadLines(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	defer file.Close()
	var lines []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lines = append(lines, line)
	}
	return lines
}

