package combgen

func CalculateCombinationsCount(input []string, maxAllowedChars int, charRepeat bool) int {
	combinationStruct := buildCombinationStruct(input, maxAllowedChars, charRepeat)
	var result int
	if !charRepeat {
		combInterface :=  nonRepeatCombinations{combinationStruct}
		result =  combInterface.calculateNrOfCombinations() 
	} else {
		combInterface :=  charRepeatCombinations{combinationStruct}
		result =  combInterface.calculateNrOfCombinations()
	}	
	
	return result
}

func CalculateCombinations(input []string, maxAllowedChars int, charRepeat bool) []string {
	combinationStruct := buildCombinationStruct(input, maxAllowedChars, charRepeat)
	var result []string
	if !charRepeat {
		combInterface :=  nonRepeatCombinations{combinationStruct}
		result = combInterface.calculateCombinations()
	} else {
		combInterface :=  charRepeatCombinations{combinationStruct}
		result =  combInterface.calculateCombinations()
	}	

	return result
}




// TODO functionality
func GetCombinationsCharRepeat(input []string, maxAllowedChars int) {
}
func GetNrOfCombinationsCharRepeat(input []string, maxAllowedChars int) {
}
