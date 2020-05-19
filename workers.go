package combgen

// Function for getting the factorial of integer value.
func getFactorial(nrOfChars int) int {
	factorial := 1;
	for i := 1; i <= nrOfChars; i++ {
		factorial = i * factorial
	}

	return factorial
}

// Function for performing the permutations for the combination list calculation.
func permute(combPointer *nonRepeatCombinations, prevLockedEl string, activeSlice[] string) {
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


