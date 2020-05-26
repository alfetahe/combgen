package combgen

func CalculateCombinationsCount(input []string, maxAllowedChars int, charRepeat bool) int {
	combInterface := getCombInterface(input, maxAllowedChars, charRepeat)

	return combInterface.calculateNrOfCombinations()
}

func CalculateCombinations(input []string, maxAllowedChars int, charRepeat bool) []string {
	combInterface := getCombInterface(input, maxAllowedChars, charRepeat)

	return combInterface.calculateCombinations()
}

func getCombInterface(input []string, maxAllowedChars int, charRepeat bool) combgenInterface {
	combinationStruct := buildCombinationStruct(input, maxAllowedChars, charRepeat)

	if !charRepeat {
		return nonRepeatCombinations{combinationStruct}
	}

	return charRepeatCombinations{combinationStruct}
}
