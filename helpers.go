package combgen

// HELPER FUNCTIONS

// RemoveFromSliceByIndex Removes item from slice by index.
func RemoveFromSliceByIndex(slice []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, slice[:index]...)
	return append(ret, slice[index+1:]...)
}

// Equal tells whether a and b contain the same elements.
func Equal(a[]string, b[]string) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}