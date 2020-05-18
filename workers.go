package combgen

// Function for getting the factorial of integer value.
func getFactorial(nrOfChars int) int {
	factorial := 1;
	for i := 1; i <= nrOfChars; i++ {
		factorial = i * factorial
	}

	return factorial
}

// TODO remove,append and return
func permute(combPointer nonRepeatCombinations, prevLockedEl string) nonRepeatCombinations {
	if combPointer.maxAllowedCharacters == len(prevLockedEl) {
		combPointer.combinationList = append(combPointer.combinationList, prevLockedEl)
	} else {
		for index, lockedEl := range combPointer.activeSlice {
			combPointer.activeSlice = RemoveFromSliceByIndex(combPointer.activeSlice, index)

			permute(combPointer, prevLockedEl+lockedEl)
		}
	}

	return combPointer
}


