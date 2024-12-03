package utils

func RemoveIndex(s []string, index int) []string {
	newStr := make([]string, len(s))
	copy(newStr, s)

	// Check if the index is valid
	if index < 0 || index >= len(newStr) {
		return newStr // Return original slice if index is out of bounds
	}
	// Append elements before and after the specified index
	return append(newStr[:index], newStr[index+1:]...)
}
