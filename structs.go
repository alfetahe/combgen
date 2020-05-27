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
	var defaultElement = ""
	permute(&noRepComb, defaultElement, noRepComb.clientInput)
	return noRepComb.combinationList
}

// TODO functionality
func (charRepComb charRepeatCombinations) calculateCombinations() []string {

	createDuplicatesForRepeat(&charRepComb)

	var defaultElement string
	permuteWithRepeat(&charRepComb, defaultElement, charRepComb.clientInput)
	return charRepComb.combinationList
}
