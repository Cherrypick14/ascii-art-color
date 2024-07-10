package ascii

import (
	"fmt"
	"os"
	"strings"

	"color/colors"
)

// ANSI color codes
var colorMap = map[string]string{
	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
	"orange":  "\033[38;5;202m",
	"pink":    "\033[38;5;198m",
	"reset":   "\033[0m",
}

var (
	inputString string
	subString   string
	colorCode   string
)

// ProcessColorFlag to process diffent flag notations that are passed as argument.
func ProcessColorFlag(colorFlag string, args []string) {
	// Case for: go run . --color=<color> <subString> <string>
	colorFlag = strings.ToLower(colorFlag)
	if len(args) == 2 {
		inputString = args[1]
		subString = args[1]
	} else if len(args) == 3 {
		subString = args[1]
		inputString = args[2]
	} else {
		PrintUsage()
		return
	}

	var code int
	// handle HEX flag notation eg #ff0000.
	if strings.HasPrefix(colorFlag, "#") {
		code = colors.HexTo256ColorCode(colorFlag)
		colorMap[colorFlag] = "\033[38;5;" + fmt.Sprint(code) + "m"
		// Handle RGB flag notation e.g. rgb(255, 0, 0)
	} else if strings.HasPrefix(colorFlag, "rgb") {
		code = colors.RgbTo256ColorCode(colorFlag)
		colorMap[colorFlag] = "\033[38;5;" + fmt.Sprint(code) + "m"
	}

	// Check if the provided color is valid
	var ok bool
	colorCode, ok = colorMap[colorFlag]
	if !ok {
		fmt.Printf("Unknown color: %s\n", colorFlag)
		os.Exit(0)
	}
}
