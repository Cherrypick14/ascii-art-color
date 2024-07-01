package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// ANSI color codes
var colors = map[string]string{
	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
	"reset":   "\033[0m",
}

func main() {
	// Define the color flag
	colorFlag := flag.String("color", "", "color for the substring")

	// Parse flags
	flag.Parse()

	// Check if enough arguments are provided
	if *colorFlag == "" || len(flag.Args()) < 1 {
		fmt.Println("Usage: go run . --color=<color> <substring to be colored> [STRING]")
		return
	}

	// Extract arguments
	args := flag.Args()
	var substring, inputString string

	if len(args) == 1 {
		substring = ""
		inputString = args[0]
	} else {
		substring = args[0]
		inputString = strings.Join(args[1:], " ")
	}

	// Get the color code
	colorCode, ok := colors[*colorFlag]
	if !ok {
		fmt.Printf("Unknown color: %s\n", *colorFlag)
		return
	}
	// debug the color code and the substring and the input string

	// fmt.Printf("Colorcode is %s\n", colorCode)
	// fmt.Printf("Substring is %s\n", substring)
	// fmt.Printf("Inputstring is %s\n", inputString)

	// Read the default banner file (standard)
	banner, err := readBannerFile("standard")
	if err != nil {
		fmt.Printf("Error reading banner file: %v\n", err)
		return
	}

	// Generate and color the ASCII art
	asciiArt := AsciiArt([]string{inputString}, banner, substring, colorCode)

	// Print the result
	fmt.Println(asciiArt)
}

func readBannerFile(banner string) ([]string, error) {
	fileName := banner + ".txt"
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var contents []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	return contents, scanner.Err()
}

func AsciiArt(words []string, contents2 []string, substring, colorCode string) string {
	var result strings.Builder
	reset := "\033[0m" // Reset color code

	fmt.Printf("Colorcode is  %s\n", colorCode)
	fmt.Printf("Substring is  %s\n", substring)

	countSpace := 0
	for _, word := range words {
		if word != "" {
			for i := 0; i < 8; i++ {
				for _, char := range word {
					if char == '\n' {
						continue
					}
					if !(char >= 32 && char <= 126) {
						return "Error: Input contains non-ASCII characters"
					}

					// Calculate the index of 'char' in the ASCII art content2.
					index := int(char-' ')*9 + 1 + i
					line := contents2[index]

					fmt.Println("Original line is", line)

					// Apply color to all substring occurrences in line
					if substring != "" && strings.Contains(line, substring) {
						// Replace all the occurrences(hence the -1) of the substring with the colored substring and then reset code
						coloredLine := strings.Replace(line, substring, colorCode+substring+reset, -1)
						fmt.Println("Colored line is", coloredLine)
						result.WriteString(coloredLine)
					} else {
						result.WriteString(line)
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

// func colorSubstring(input, substring, colorCode string) string {
// 	if substring == "" {
// 		// Color the entire string
// 		return colorCode + input + colors["reset"]
// 	}

// 	// Color only the specified substring
// 	coloredSubstring := colorCode + substring + colors["reset"]
// 	return strings.Replace(input, substring, coloredSubstring, -1)
// }

// package main

// import (
// 	"flag"
// 	"fmt"
// 	"os"
// 	"strings"
// 	"collors/utils"
// )

// func main() {
// 	colorFlag := flag.String("color", "", "Color and substring to be colored")
// 	flag.Parse()

// 	if *colorFlag == "" || flag.NArg() < 1 {
// 		printUsage()
// 		os.Exit(1)
// 	}

// 	colorParts := strings.SplitN(*colorFlag, " ", 2)
// 	color := strings.TrimPrefix(colorParts[0], "--color=")
// 	var substring string
// 	if len(colorParts) > 1 {
// 		substring = colorParts[1]
// 	}

// 	text := flag.Arg(0)

// 	ansiColor, err := parseColor(color)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		os.Exit(1)
// 	}

// 	coloredText := colorizeString(text, substring, ansiColor)
// 	newColoredText := utils.AsciiArt(coloredText)
// 	fmt.Println(newColoredText) // Replace this with your ASCII art generator function call
// }

// func parseColor(color string) (string, error) {
// 	switch strings.ToLower(color) {
// 	case "red":
// 		return "\033[31m", nil
// 	case "green":
// 		return "\033[32m", nil
// 	case "blue":
// 		return "\033[34m", nil
// 	case "yellow":
// 		return "\033[33m", nil
// 	case "purple":
// 		return "\033[35m", nil
// 	case "cyan":
// 		return "\033[36m", nil
// 	case "white":
// 		return "\033[37m", nil
// 	default:
// 		return "", fmt.Errorf("unsupported color: %s", color)
// 	}
// }

// func colorizeString(text, substring, color string) string {
// 	if substring == "" {
// 		return color + text + "\033[0m"
// 	}
// 	return strings.ReplaceAll(text, substring, color+substring+"\033[0m")
// }

// func printUsage() {
// 	fmt.Println("Usage: go run . [OPTION] [STRING]")
// 	fmt.Println("EX: go run . --color=<color> <substring to be colored> \"something\"")
// }
