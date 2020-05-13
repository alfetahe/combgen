package combgen

type combinations struct {
	clientInput [] string
	maxAllowedCharacters int
	nrOfCombinations int
	combinationList []string
	activeSlice []string
}

type nonRepeatCombinations struct {
	combinations
}

type charRepeatCombinations struct {
	combinations
}


// Get Method for calculated struct data. Return int value of nrOfPossibleCombinations.
func getNrOfCombinations(combInterf combgenInterface) int {
	return combInterf.calculateNrOfCombinations() 
}

// Get Method for calculated struct data. Return int value of nrOfPossibleCombinations.
func getCombinations(combInterf combgenInterface) []string {
	return combInterf.calculateCombinations() 
}


// Interface Calculation method.
func (noRepComb nonRepeatCombinations) calculateNrOfCombinations() int {
	nrOfChars := len(noRepComb.clientInput)	
	result := getFactorial(nrOfChars) / getFactorial(nrOfChars - noRepComb.maxAllowedCharacters)
	return result
}

// Intrface Calculation method.
func (noRepComb nonRepeatCombinations) calculateCombinations() []string {

	asp := []string{"adff"}

	var defaultElement = ""
	permute(noRepComb, defaultElement)

	return asp
}



// TODO functionality
func (charRepComb charRepeatCombinations) calculateNrOfCombinations() {
}
func (charRepComb charRepeatCombinations) calculateCombinations() {
}


