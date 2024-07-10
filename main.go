package main

import (
	"fmt"
	"os"

	"color/ascii"
)

func main() {
	if len(os.Args) < 2 {
		ascii.PrintUsage()
	}
	args := os.Args[1:]
	data := ascii.ParseFlagsAndArgs(args)

	// Split our input text to a string slice and separate with a newline.
	words := ascii.ProcessInput(data.Text)

	// Read the banner file and get its contents.
	bannerContents, err := ascii.ReadBanner(data.Banner)
	if err != nil {
		fmt.Printf("Error reading banner file: %v\n", err)
		return
	}

	// Generate and color the ASCII art
	asciiArt := ascii.AsciiArt(words, bannerContents, data.SubText, data.Color)

	fmt.Print(asciiArt)
}
