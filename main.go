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
	// Define flags using flag package
	colorFlag := flag.String("color", "", "color for the substring")

	// Parse flags
	flag.Parse()

	// Extract arguments using flag.Args()
	args := flag.Args()

	var substring, inputString, banner string
	var colorCode string

	if *colorFlag != "" {
		// Case for: go run . --color=<color> <substring> <string>
		if len(args) == 1 {
			inputString = args[0]
			substring = args[0]
		} else if len(args) == 2 {
			substring = args[0]
			inputString = args[1]
		} else {
			printUsage()
			return
		}

		// Check if the provided color is valid
		var ok bool
		colorCode, ok = colors[*colorFlag]
		if !ok {
			fmt.Printf("Unknown color: %s\n", *colorFlag)
			return
		}
	} else {
		// Handle cases without the --color flag
		switch len(args) {
		case 1:
			// Case for: go run . <string>
			inputString = args[0]
		case 2:
			// Case for: go run . <string> <banner>
			inputString = args[0]
			banner = args[1]
		// case 3:
		// 	// Case for: go run . <substring> <string> <banner>
		// 	substring = args[0]
		// 	inputString = args[1]
		// 	banner = args[2]
		default:
			printUsage()
			return
		}
	}

	// Read the banner file if specified
	var bannerContents []string
	var err error

	if banner != "" {
		bannerContents, err = readBannerFile(banner)
		if err != nil {
			fmt.Printf("Error reading banner file: %v\n", err)
			return
		}
	} else {
		// Read the default banner file (standard)
		bannerContents, err = readBannerFile("standard")
		if err != nil {
			fmt.Printf("Error reading banner file: %v\n", err)
			return
		}
	}

	// Generate and color the ASCII art
	asciiArt := AsciiArt([]string{inputString}, bannerContents, substring, colorCode)

	// Print the result
	fmt.Println(asciiArt)
}

// func main() {
// 	// Define flags using flag package
// 	colorFlag := flag.String("color", "", "color for the substring")

// 	// Parse flags
// 	flag.Parse()

// 	// Extract arguments using os package for flexibility
// 	args := flag.Args()
// 	fmt.Println(len(args))
// 	fmt.Println(args)
// 	var substring, inputString, banner string
//     var colorCode string

// 	// switch len(args) {
// 	// case 1:
// 	// 	// Case for: go run . <string>
// 	// 	inputString = args[0]
// 	// case 2:
// 	// 	// Case for: go run . <string> <banner>
// 	// 	inputString = args[0]

// 	// 	banner = args[1]

// 	// case 3:
// 	// 	substring = args[0]
// 	// 	inputString = args[1]
// 	// 	banner = args[2]
// 	// default:
// 	// 	printUsage()
// 	// 	return
// 	// }
// 	// switch len(args) {
// 	// case 1:
// 	// 	// Case for: go run . <string>
// 	// 	inputString = args[0]
// 	// 	banner = "standard" // Default banner
// 	// case 2:
// 	// 	// Case for: go run . <string> <banner>
// 	// 	substring = args[0]
// 	// 	inputString = args[1]

// 	// 	// banner = args[1]

// 	// default:
// 	// 	printUsage()
// 	// 	return
// 	// }

// 	// Initialize colorCode

// 	// Check if a color flag was provided
// 	if *colorFlag != "" {
// 		colorCode = colors[*colorFlag]
// 		if colorCode == "" {
// 			fmt.Printf("Unknown color: %s\n", *colorFlag)
// 			return
// 		}
// 		// If color is provided, use the first word as substring
// 		words := strings.Fields(inputString)
// 		if len(words) > 0 {
// 			substring = words[0]
// 		}
// 	}

// 	// Read the banner file if specified
// 	var bannerContents []string
// 	if banner != "" {
// 		var err error
// 		bannerContents, err = readBannerFile(banner)
// 		if err != nil {
// 			fmt.Printf("Error reading banner file: %v\n", err)
// 			return
// 		}
// 	} else {
// 		// Read the default banner file (standard)
// 		var err error
// 		bannerContents, err = readBannerFile("standard")
// 		if err != nil {
// 			fmt.Printf("Error reading banner file: %v\n", err)
// 			return
// 		}
// 	}

// 	// Generate and color the ASCII art
// 	asciiArt := AsciiArt([]string{inputString}, bannerContents, substring, colorCode)

// 	// Print the result
// 	fmt.Println(asciiArt)
// }

// func main() {
// 	// Define flags using flag package
// 	colorFlag := flag.String("color", "", "color for the substring")

// 	// Parse flags
// 	flag.Parse()

// 	// Extract arguments using os package for flexibility
// 	args := flag.Args()
// 	var substring, inputString, banner string

// 	switch len(args) {
// 	case 0:
// 		printUsage()
// 		return
// 	case 1:
// 		substring = ""
// 		inputString = args[0]
// 	case 2:
// 		substring = args[0]
// 		inputString = args[1]
// 	case 3:
// 		substring = args[0]
// 		inputString = args[1]
// 		banner = args[2]
// 	default:
// 		printUsage()
// 		return
// 	}

// 	// Initialize colorCode

// 	var colorCode string

// 	// Check if a color flag was provided

// 	if *colorFlag != "" {
// 		colorCode = colors[*colorFlag]
// 		if colorCode == "" {
// 			fmt.Printf("Unknown color: %s\n", *colorFlag)
// 			return
// 		}
// 	}

// 	// Read the banner file if specified
// 	var bannerContents []string
// 	if banner != "" {
// 		var err error
// 		bannerContents, err = readBannerFile(banner)
// 		if err != nil {
// 			fmt.Printf("Error reading banner file: %v\n", err)
// 			return
// 		}
// 	} else {
// 		// Read the default banner file (standard)
// 		var err error
// 		bannerContents, err = readBannerFile("standard")
// 		if err != nil {
// 			fmt.Printf("Error reading banner file: %v\n", err)
// 			return
// 		}
// 	}

// 	// Generate and color the ASCII art
// 	asciiArt := AsciiArt([]string{inputString}, bannerContents, substring, colorCode)

// 	// Print the result
// 	fmt.Println(asciiArt)

// }
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

	countSpace := 0
	sub := false
	position := 0

	for _, word := range words {
		indexs := subIndexs(word, substring)
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
						return "Error: Input contains non-ASCII characters"
					}
					index := int(char-' ')*9 + 1 + i
					result.WriteString(contents2[index])

					if sub && k == position+len(substring)-1 {
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

func subIndexs(s, subStr string) []int {
	index := []int{}
	leftCharacters := 0

	for len(subStr) > 0 {
		idx := strings.Index(s, subStr)
		if idx == -1 {
			break
		}
		index = append(index, idx+leftCharacters)
		s = s[idx+len(subStr):]
		leftCharacters += idx + len(subStr)
	}
	return index
}

func validIndex(index int, indexs []int) bool {
	for _, idx := range indexs {
		if index == idx {
			return true
		}
	}
	return false
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  To colorize a substring in ASCII art:")
	fmt.Println("    go run . --color=<color> <substring to be colored>")
	fmt.Println("  To specify a custom banner file:")
	fmt.Println("    go run . <substring to be colored> <banner>")
	fmt.Println("  Example:")
	fmt.Println("    go run . --color=red hello")
	fmt.Println("    go run . hello standard")
}

// func main() {
// 	// Define the color flag
// 	colorFlag := flag.String("color", "", "color for the substring")

// 	// Parse flags
// 	flag.Parse()

// 	// Check if enough arguments are provided
// 	if *colorFlag == "" || len(flag.Args()) < 1 {
// 		printUsage()
// 		return
// 	}

// 	// Extract arguments
// 	args := flag.Args()
// 	var substring, inputString string

// 	if len(args) == 1 {
// 		substring = ""
// 		inputString = args[0]
// 	} else if len(args) == 2 {
// 		substring = args[0]
// 		inputString = args[1]
// 		//inputString = strings.Join(args[1:], " ")
// 	}

// 	// Get the color code
// 	colorCode, ok := colors[*colorFlag]
// 	if !ok {
// 		fmt.Printf("Unknown color: %s\n", *colorFlag)
// 		return
// 	}

// 	// Read the default banner file (standard)
// 	banner, err := readBannerFile("standard")
// 	if err != nil {
// 		fmt.Printf("Error reading banner file: %v\n", err)
// 		return
// 	}

// 	// Generate and color the ASCII art
// 	asciiArt := AsciiArt([]string{inputString}, banner, substring, colorCode)

// 	// Print the result
// 	fmt.Println(asciiArt)
// }

// func readBannerFile(banner string) ([]string, error) {
// 	fileName := banner + ".txt"
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	var contents []string
// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {
// 		contents = append(contents, scanner.Text())
// 	}

// 	return contents, scanner.Err()
// }

// func AsciiArt(words []string, contents2 []string, substring, colorCode string) string {
// 	var result strings.Builder
// 	reset := "\033[0m" // Reset color code

// 	fmt.Printf("Colorcode is  %s\n", colorCode)
// 	fmt.Printf("Substring is  %s\n", substring)

// 	countSpace := 0
// 	sub := false // checker
// 	position := 0

// 	for _, word := range words {
// 		indexs := subIndexs(word, substring)
// 		if word != "" {
// 			for i := 0; i < 8; i++ {
// 				for k, char := range word {
// 					if validIndex(k, indexs) {
// 						result.WriteString(colorCode)
// 						sub = true
// 						position = k

// 					}
// 					if char == '\n' {
// 						continue
// 					}
// 					if !(char >= 32 && char <= 126) {
// 						return "Error: Input contains non-ASCII characters"
// 					}
// 					// Calculate the index of 'char' in the ASCII art content2.
// 					index := int(char-' ')*9 + 1 + i

// 					// Append the corresponding ASCII art line for the character.
// 					result.WriteString(contents2[index])

// 					if sub && k == position+len(substring)-1 {
// 						result.WriteString(reset)
// 					}

// 				}
// 				result.WriteString("\n")
// 			}
// 		} else {
// 			countSpace++
// 			if countSpace < len(words) {
// 				result.WriteString("\n")
// 			}
// 		}
// 	}
// 	return result.String()
// }

// func subIndexs(s, subStr string) []int {
// 	index := []int{}
// 	leftCharacters := 0

// 	for len(subStr) > 0 {
// 		idx := strings.Index(s, subStr)
// 		if idx == -1 {
// 			break
// 		}
// 		index = append(index, idx+leftCharacters)
// 		s = s[idx+len(subStr):]
// 		leftCharacters += idx + len(subStr)
// 	}
// 	return index
// }

// func validIndex(index int, indexs []int) bool {
// 	for _, idx := range indexs {
// 		if index == idx {
// 			return true
// 		}
// 	}
// 	return false
// }

// func printUsage() {
// 	fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\"")

// }

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
