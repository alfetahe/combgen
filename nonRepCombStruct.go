package combgen

type nonRepeatCombinations struct {
	combinations
}

func (noRepComb nonRepeatCombinations) calculateNrOfCombinations() int {
	nrOfChars := len(noRepComb.clientInput)
	result := getFactorial(nrOfChars) / getFactorial(nrOfChars-noRepComb.maxAllowedCharacters)
	return result
}

func (noRepComb nonRepeatCombinations) calculateCombinations() []string {
	noRepComb.permute("", noRepComb.clientInput)
	return noRepComb.combinationList
}

func (combPointer *nonRepeatCombinations) permute(prevLockedEl string, activeSlice []string) {
	if combPointer.maxAllowedCharacters == len(prevLockedEl) {
		combPointer.combinationList = append(combPointer.combinationList, prevLockedEl)
	} else {
		reduceCharacaters(combPointer, activeSlice, prevLockedEl)
	}
}
