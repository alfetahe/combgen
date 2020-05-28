package combgen

type charRepeatCombinations struct {
	combinations
}

func (charRepComb charRepeatCombinations) calculateNrOfCombinations() int {
	nrOfChars := len(charRepComb.clientInput)
	result := calculateExponent(nrOfChars, charRepComb.maxAllowedCharacters)
	return result
}

func (charRepComb charRepeatCombinations) calculateCombinations() []string {
	charRepComb.createDuplicatesForRepeat()
	charRepComb.permute("", charRepComb.clientInput)
	return charRepComb.combinationList
}

func (combPointer *charRepeatCombinations) permute(prevLockedEl string, activeSlice []string) {
	if combPointer.maxAllowedCharacters == len(prevLockedEl) {
		if !sliceContains(combPointer.combinationList, prevLockedEl) {
			combPointer.combinationList = append(combPointer.combinationList, prevLockedEl)
		}

		return
	}
	reduceCharacaters(combPointer, activeSlice, prevLockedEl)
}

func (combPointer *charRepeatCombinations) createDuplicatesForRepeat() {
	for _, element := range combPointer.clientInput {
		for i := 1; i < combPointer.maxAllowedCharacters; i++ {
			combPointer.clientInput = append(combPointer.clientInput, element)
		}
	}
}
