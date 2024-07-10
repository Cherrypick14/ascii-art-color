package ascii

import (
	"flag"
	"strings"
)

var banner string

// ParseFlagsAndFlags processes commandline arguments and returns a Data.
func ParseFlagsAndArgs(args []string) Data {
	// Define flags using flag package
	colorFlag := flag.String("color", "", "color for the subString")
	// Check for undefined flags.
	if strings.HasPrefix(args[0], "-") && !strings.HasPrefix(args[0], "--color=") {
		PrintUsage()
	}
	// Parse flags
	flag.Parse()

	// Check if --color flag is provided and validate it
	if *colorFlag != "" && strings.HasPrefix(args[0], "--color=") {
		ProcessColorFlag(*colorFlag, args)
	} else if !strings.HasPrefix(args[0], "--color=") && len(args) >= 1 {
		HandleArguments(args)
	} else {
		PrintUsage()
	}
	return Data{
		Text:    inputString,
		Banner:  banner,
		SubText: subString,
		Color:   colorCode,
	}
}

// HandleArguments handles arguments passed	without the --color flag.
func HandleArguments(args []string) Data {
	switch len(args) {
	case 1:
		// Case for: go run . <string>
		inputString = args[0]
	case 2:
		// Case for: go run . <string> <banner>
		inputString = args[0]
		banner = args[1]
	default:
		PrintUsage()
	}
	return Data{
		Text:    inputString,
		Banner:  banner,
		SubText: subString,
		Color:   colorCode,
	}
}
