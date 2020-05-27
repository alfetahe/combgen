package combgen

type combinations struct {
	clientInput          []string
	maxAllowedCharacters int
	nrOfCombinations     int
	combinationList      []string
}

func (comb *combinations) setMaxAllowedCharSize(charRepeat bool) {
	maxAllowedCharSize := comb.maxAllowedCharacters
	if maxAllowedCharSize == 0 || maxAllowedCharSize > len(comb.clientInput) && !charRepeat {
		comb.maxAllowedCharacters = len(comb.clientInput)
	} else {
		comb.maxAllowedCharacters = maxAllowedCharSize
	}
}
