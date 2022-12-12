package util

import (
	"os"
	"strings"
)

// MustReadInput reads input and parse it by lines
func MustReadInput() []string {
	fileData, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileData), "\n")

	if lines[len(lines)-1] == "" { // ignore last empty line if exist
		lines = lines[:len(lines)-1]
	}

	return lines
}

func MustReadTestInput() []string {
	fileData, err := os.ReadFile("input.test.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileData), "\n")

	if lines[len(lines)-1] == "" { // ignore last empty line if exist
		lines = lines[:len(lines)-1]
	}

	return lines
}
