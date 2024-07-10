package ascii

import "strings"

// SubIndexs finds all occurrences of subStr in inputStr and returns their indices.
func subIndexs(inputStr, subStr string) []int {
	index := []int{}
	leftCharacters := 0

	for len(subStr) > 0 {
		idx := strings.Index(inputStr, subStr)
		if idx == -1 {
			break
		}
		index = append(index, idx+leftCharacters)
		inputStr = inputStr[idx+len(subStr):] // update inputStr after substring ocurrence.
		leftCharacters += idx + len(subStr)
	}
	return index
}

// ValidIndex checks if index exits in the indexs slice.
func validIndex(index int, indexs []int) bool {
	for _, idx := range indexs {
		if index == idx {
			return true
		}
	}
	return false
}
