package combgen

// HELPER FUNCTIONS

// RemoveFromSliceByIndex Removes item from slice by index.
func RemoveFromSliceByIndex(slice []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, slice[:index]...)
	ret = append(ret, slice[index+1:]...)
	return ret
}

// SliceContains takes a slice and looks for an element in it. If found it will
// return true, otherwise false.
func SliceContains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
