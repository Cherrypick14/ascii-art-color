package ascii

import (
	"fmt"
	"os"
	"strings"
)

// ProcessInput processes the input string and returns a slice of words.
func ProcessInput(inputString string) []string {
	switch inputString {
	case "":
		os.Exit(0)
	case "\\a", "\\0", "\\f", "\\v", "\\r":
		fmt.Println("Error: Non printable character", inputString)
		os.Exit(0)
	}

	inputString = strings.ReplaceAll(inputString, "\\t", "    ")
	inputString = strings.ReplaceAll(inputString, "\\b", "\b")
	inputString = strings.ReplaceAll(inputString, "\\n", "\n")

	// Process backspaces (\b)
	for i := 0; i < len(inputString); i++ {
		indexB := strings.Index(inputString, "\b")
		if indexB > 0 {
			inputString = inputString[:indexB-1] + inputString[indexB+1:]
		}
	}
	// Split our input text to a string slice and separate with a newline.
	words := strings.Split(inputString, "\n")

	return words
}
