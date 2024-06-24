package main

import (
  "flag"
  "fmt"
  "os"
  "strings"

  "collors/utils/colors"
)
// Existing ASCII art generator function (replace contents with your actual data)
func AsciiArt(words []string, contents2 []string) string {
	var result strings.Builder
  
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
			// Print the calculated index of 'char' Ascii Art in content2.
			result.WriteString(contents2[int(char-' ')*9+1+i])
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

  func colorizeArt(art, color, substring string) string {
	coloredString := colors.GetColor(color)
  
	if substring == "" {
	  return coloredString + art + colors.reset
	} else {
	  startIndex := strings.Index(art, substring)
	  if startIndex == -1 {
		fmt.Println("Error: substring not found in the ASCII art")
		os.Exit(1)
	  }
	  endIndex := startIndex + len(substring)
	  return coloredString + art[:startIndex] + substring + colors.reset + art[endIndex:]
	}
  }
  
  func main() {
	colorFlag := flag.String("color", "", "color to apply (red, green, yellow, blue, magenta, cyan) [optional:substring to color]")
	text := flag.Arg(0) // Assuming the first argument is the text
  
	flag.Parse()
  
	colorAndSubstr := *colorFlag
	parts := strings.SplitN(colorAndSubstr, " ", 2)
  
	if len(parts) == 0 || len(parts) > 2 {
	  fmt.Println("Error: invalid color flag format")
	  flag.PrintDefaults()
	  os.Exit(1)
	}
  
	color := parts[0]
	var substring string
	if len(parts) == 2 {
	  substring = parts[1]
	}
  
	art := AsciiArt(strings.Split(text, ""), /* replace with your actual ASCII art data */)
	coloredArt := colorizeArt(art, color, substring)
  
	fmt.Println(coloredArt)
  }
