package functions

// NewLines checks if the provided slice of strings contains only newlines.
func NewLines(lines []string) string {
	newL := ""
	for i, line := range lines {
		if line != "" {
			return "false"
		}
		if i == 0 {
			continue
		}
		newL += "\n"
	}
	return newL
}
