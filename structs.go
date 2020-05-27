package combgen

type combinations struct {
	clientInput          []string
	maxAllowedCharacters int
	nrOfCombinations     int
	combinationList      []string
}

type nonRepeatCombinations struct {
	combinations
}

type charRepeatCombinations struct {
	combinations
}

func (comb *combinations) setMaxAllowedCharSize(charRepeat bool) {
	maxAllowedCharSize := comb.maxAllowedCharacters
	if maxAllowedCharSize == 0 || maxAllowedCharSize > len(comb.clientInput) && !charRepeat {
		comb.maxAllowedCharacters = len(comb.clientInput)
	} else {
		comb.maxAllowedCharacters = maxAllowedCharSize
	}
}

func (noRepComb nonRepeatCombinations) calculateNrOfCombinations() int {
	nrOfChars := len(noRepComb.clientInput)
	result := getFactorial(nrOfChars) / getFactorial(nrOfChars-noRepComb.maxAllowedCharacters)
	return result
}

func (charRepComb charRepeatCombinations) calculateNrOfCombinations() int {
	nrOfChars := len(charRepComb.clientInput)
	result := calculateExponent(nrOfChars, charRepComb.maxAllowedCharacters)
	return result
}

func (noRepComb nonRepeatCombinations) calculateCombinations() []string {
	noRepComb.permute(&noRepComb, "", noRepComb.clientInput)
	return noRepComb.combinationList
}

func (charRepComb charRepeatCombinations) calculateCombinations() []string {
	createDuplicatesForRepeat(&charRepComb)
	charRepComb.permute(&charRepComb, "", charRepComb.clientInput)
	return charRepComb.combinationList
}

func (charRepComb *charRepeatCombinations) permute(combPointer *charRepeatCombinations, prevLockedEl string, activeSlice []string) {
	if combPointer.maxAllowedCharacters == len(prevLockedEl) {
		if !SliceContains(combPointer.combinationList, prevLockedEl) {
			combPointer.combinationList = append(combPointer.combinationList, prevLockedEl)
		}

		return
	}

	for index, lockedEl := range activeSlice {
		newActiveSlice := RemoveFromSliceByIndex(activeSlice, index)
		newLockedEl := prevLockedEl + lockedEl
		combPointer.permute(combPointer, newLockedEl, newActiveSlice)
	}
}

func (noRepComb *nonRepeatCombinations) permute(combPointer *nonRepeatCombinations, prevLockedEl string, activeSlice []string) {
	if combPointer.maxAllowedCharacters == len(prevLockedEl) {
		combPointer.combinationList = append(combPointer.combinationList, prevLockedEl)
	} else {
		for index, lockedEl := range activeSlice {
			newActiveSlice := RemoveFromSliceByIndex(activeSlice, index)
			newLockedEl := prevLockedEl + lockedEl
			combPointer.permute(combPointer, newLockedEl, newActiveSlice)
		}
	}
}
