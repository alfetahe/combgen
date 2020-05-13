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
func permute(combPointer nonRepeatCombinations, lockedElement string) {
	if combPointer.maxAllowedCharacters == len(lockedElement) {
		//append(combPointer.combinationList, lockedElement)
	} else {
		for index, lockedEl := range combPointer.activeSlice {
			RemoveFromSliceByIndex(combPointer.activeSlice, index)

			permute(combPointer, lockedElement+lockedEl)
		}
	}

	//return combPointer.combinationList
}


