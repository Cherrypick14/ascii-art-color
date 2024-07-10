package ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadBanner(banner string) ([]string, error) {
	// Check if banner is defined; if not, use default "standard"
	if banner == "" {
		banner = "standard"
	}

	fileName := strings.ToLower(banner)
	if !strings.HasSuffix(fileName, ".txt") {
		fileName += ".txt"
	}

	// Confirm file information.
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return nil, fmt.Errorf("error reading file information: %w", err)
	}
	fileSize := fileInfo.Size()

	// Check if file size is one of the expected sizes
	if fileSize != 6623 && fileSize != 4702 && fileSize != 7462 {
		return nil, fmt.Errorf("error with the file size: %d", fileSize)
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var contents []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file: %w", err)
	}

	return contents, nil
}
