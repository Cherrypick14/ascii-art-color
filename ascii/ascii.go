package ascii

import (
	"strings"
)

// AsciiArt processes words, printing their ASCII art
// character by character and adding new lines as needed.
func AsciiArt(words []string, contents2 []string, subString, colorCode string) string {
	var result strings.Builder
	reset := "\033[0m" // Reset color code

	countSpace := 0
	sub := false
	position := 0

	for _, word := range words {
		indexs := subIndexs(word, subString)
		if word != "" {
			for i := 0; i < 8; i++ {
				for k, char := range word {
					if validIndex(k, indexs) && colorCode != "" {
						result.WriteString(colorCode)
						sub = true
						position = k
					}
					if char == '\n' {
						continue
					}
					if !(char >= 32 && char <= 126) {
						return "Error: Non printable ASCII character\n"
					}
					index := int(char-' ')*9 + 1 + i
					result.WriteString(contents2[index])

					if sub && k == position+len(subString)-1 {
						result.WriteString(reset)
					}
				}
				result.WriteString("\n")
			}
		} else {
			countSpace++
			if countSpace < len(words) {
				result.WriteString("\n")
			}
		}
	}
	return result.String()
}
