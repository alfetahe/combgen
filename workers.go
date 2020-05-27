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

// Function for performing the permutations for the combination list calculation.
func permute(combPointer *nonRepeatCombinations, prevLockedEl string, activeSlice []string) {
	if combPointer.maxAllowedCharacters == len(prevLockedEl) {
		combPointer.combinationList = append(combPointer.combinationList, prevLockedEl)
	} else {
		for index, lockedEl := range activeSlice {
			newActiveSlice := RemoveFromSliceByIndex(activeSlice, index)
			newLockedEl := prevLockedEl + lockedEl
			permute(combPointer, newLockedEl, newActiveSlice)
		}
	}
}

func permuteWithRepeat(combPointer *charRepeatCombinations, prevLockedEl string, activeSlice []string) {
	if combPointer.maxAllowedCharacters == len(prevLockedEl) {
		if !SliceContains(combPointer.combinationList, prevLockedEl) {
			combPointer.combinationList = append(combPointer.combinationList, prevLockedEl)
		}

		return
	}

	for index, lockedEl := range activeSlice {
		newActiveSlice := RemoveFromSliceByIndex(activeSlice, index)
		newLockedEl := prevLockedEl + lockedEl
		permuteWithRepeat(combPointer, newLockedEl, newActiveSlice)
	}

}

func createDuplicatesForRepeat(combPointer *charRepeatCombinations) {
	for _, element := range combPointer.clientInput {
		for i := 1; i < combPointer.maxAllowedCharacters; i++ {
			combPointer.clientInput = append(combPointer.clientInput, element)
		}
	}
}
