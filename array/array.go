package array

// Contains returns true if an item exists in the array, false otherwise
func Contains(haystack []string, needle string) bool {
	for _, straw := range haystack {
		if straw == needle {
			return true
		}
	}
	return false
}

// ContainsRune returns true if an item exists in the array, false otherwise
func ContainsRune(haystack []rune, needle rune) bool {
	for _, straw := range haystack {
		if straw == needle {
			return true
		}
	}
	return false
}

// Prepend ...
func Prepend(array []string, items ...string) []string {
	return append(items, array...)
}
