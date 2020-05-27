package combgen

// Function for getting the factorial of integer value.
func getFactorial(nrOfChars int) int {
	factorial := 1
	for i := 1; i <= nrOfChars; i++ {
		factorial = i * factorial
	}

	return factorial
}

// Function for calculating the exponent
func calculateExponent(nrOfChars int, maxAllowedCharacters int) int {
	result := nrOfChars
	for i := 1; i < maxAllowedCharacters; i++ {
		result = nrOfChars * result
	}

	return result
}

func createDuplicatesForRepeat(combPointer *charRepeatCombinations) {
	for _, element := range combPointer.clientInput {
		for i := 1; i < combPointer.maxAllowedCharacters; i++ {
			combPointer.clientInput = append(combPointer.clientInput, element)
		}
	}
}

func reduceCharacaters(combPointer combgenInterface) {

}
