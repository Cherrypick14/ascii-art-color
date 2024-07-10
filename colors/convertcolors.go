package colors

import (
	"fmt"
	"os"
	"strings"
)

// Convert RGB color values to 256-color value.
func RgbTo256ColorCode(rgb string) int {
	var r, g, b int
	rgb = strings.TrimPrefix(rgb, "rgb(")
	rgb = strings.TrimSuffix(rgb, ")")
	// scan through values to convert to code point values.
	_, err := fmt.Sscanf(rgb, "%d,%d,%d", &r, &g, &b)
	if err != nil || r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 {
		fmt.Println("Invalid RGB color code: ", rgb)
		os.Exit(1)
	}
	return rgbTo256ColorValues(r, g, b)
}

// Convert HEX colorcode to 256-color code
func HexTo256ColorCode(hex string) int {
	hex = strings.TrimPrefix(hex, "#")
	var r, g, b int
	// scan through hex values to convert to code point values.
	_, err := fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b)
	if err != nil {
		fmt.Println("Invalid HEX color code: ", hex)
		os.Exit(1)
	}
	return rgbTo256ColorValues(r, g, b)
}

// Map RGB values to 256-color code points.
func rgbTo256ColorValues(r, g, b int) int {
	var rIndex, gIndex, bIndex int = r / 51, g / 51, b / 51
	return 36*rIndex + 6*gIndex + bIndex + 16
}
