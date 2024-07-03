package functions

import (
	"os"
	"strings"
)

// ReadFromFile gets the filename and splits it into an array of strings.
func Readfile(fileName string) []string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil
	}
	if string(file) == "" {
		return nil
	}

	var lines []string
	if fileName == "thinkertoy.txt" {
		lines = strings.Split(string(file), "\r\n")

	} else {
	lines = strings.Split(string(file), "\n")
	}
	return lines
}
