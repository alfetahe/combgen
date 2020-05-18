package combgen

func getFactorial(nrOfChars int) int {
	factorial := 1;
	for i := 1; i <= nrOfChars; i++ {
		factorial = i * factorial
	}

	return factorial
}

// Recursive function for looping thru the slice.
func permute(slice []string, prevLockedEl string) {
	if outputMaxLength == len(prevLockedEl) {
		generatedData = append(generatedData, prevLockedEl)
	} else {
		for index, lockedEl := range slice {
			newSlice := RemoveFromSliceByIndex(slice, index)
			permute(newSlice, prevLockedEl+lockedEl)
		}
	}
}

// Set max allowed string length.
func setMaxAllowedStringLength(allowedLength int, inputLength int) {
	result := checkIfAllowedStringLengthIsValid(allowedLength, inputLength)
	if result == true {
		outputMaxLength = allowedLength
	} else {
		outputMaxLength = inputLength
	}
}

func checkIfAllowedStringLengthIsValid(allowedLength int, inputLength int) bool {
	if allowedLength == 0 || allowedLength > inputLength {
		return false
	} else {
		return true
	}
}

func resetGlobalVariables() {
	generatedData = []string{}
	outputMaxLength = 0
}