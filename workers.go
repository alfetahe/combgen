package combgen

// Function for getting the factorial of integer value.
func getFactorial(nrOfChars int) int {
	factorial := 1
	for i := 1; i <= nrOfChars; i++ {
		factorial = i * factorial
	}

	return factorial
}

// Function for calculating the exponent
func calculateExponent(nrOfChars int, maxAllowedCharacters int) int {
	result := nrOfChars
	for i := 1; i < maxAllowedCharacters; i++ {
		result = nrOfChars * result
	}

	return result
}

// Function for reducing the slice length and keep the recursion going.
func reduceCharacaters(combPointer combgenInterface, activeSlice []string, prevLockedEl string) {
	for index, lockedEl := range activeSlice {
		newActiveSlice := RemoveFromSliceByIndex(activeSlice, index)
		newLockedEl := prevLockedEl + lockedEl
		combPointer.permute(newLockedEl, newActiveSlice)
	}
}
