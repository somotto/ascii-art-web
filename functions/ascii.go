package functions

import (
	"strings"
)

// AsciiArt generates ASCII art for the given words.
func AsciiArt(input string, inputFile []string) string {
	var result strings.Builder
	var newLinesOnly strings.Builder
	// Replace special characters
	input = strings.ReplaceAll(input, "\r\n", "\\n")
	input = strings.ReplaceAll(input, "\\t", "    ")

	

	// Split input by newline
	lines := strings.Split(input, "\\n")
	// Handle special case: only new lines
	if newLines := NewLines(lines); newLines != "false" {
		newLinesOnly.WriteString(newLines)
		return newLinesOnly.String()
	} else {
		for _, line := range lines {
			if line == "" {
				result.WriteString("\n")
			} else {
				for i := 0; i < len(line); {
					for j := 0; j < 8; {
						start := (int(line[i]-32) * 9) + 1
						result.WriteString(inputFile[start+j])
						i++
						if i == len(line) {
							if j == 7 {
								result.WriteString("\n")
								break
							}
							result.WriteString("\n")
							j++
							i = 0

						}
					}
				}
			}
		}
	}
	return result.String()
}
