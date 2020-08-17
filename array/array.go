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
