package combgen

// HELPER FUNCTIONS

// removeFromSliceByIndex Removes item from slice by index.
func removeFromSliceByIndex(slice []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, slice[:index]...)
	ret = append(ret, slice[index+1:]...)
	return ret
}

// sliceContains takes a slice and looks for an element in it. If found it will
// return true, otherwise false.
func sliceContains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
