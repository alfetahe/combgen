package combgen

func GetNrOfCombinationsNonRepeat(input []string, maxAllowedChars int) int {
	combs := nonRepeatCombinations{combinations{clientInput: input, maxAllowedCharacters: maxAllowedChars}}
	return getNrOfCombinations(combs)
}

func GetCombinationsNonRepeat(input []string, maxAllowedChars int) []string {
	combs := nonRepeatCombinations{
		combinations{
			clientInput: input,
			maxAllowedCharacters: maxAllowedChars,
			combinationList: input,
			activeSlice: input,
		},
	}
	
	return getCombinations(combs)
}




// TODO functionality
func GetCombinationsCharRepeat(input []string, maxAllowedChars int) {
}
func GetNrOfCombinationsCharRepeat(input []string, maxAllowedChars int) {
}
