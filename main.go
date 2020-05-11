package combgen

// Global variable for storing the generated data.
var generatedData []string

// Global variable for storing the output string max length.
var outputMaxLength int

// NrOfPossibleCalculations return the possible nr of combinations.
func NrOfPossibleCalculations(inputSlice[]string, allowedLength int) int {
	resetGlobalVariables()
	nrOfChars := len(inputSlice)	
	validCheck := checkIfAllowedStringLengthIsValid(allowedLength, nrOfChars)

	if validCheck == true {
		return getFactorial(nrOfChars) / getFactorial(nrOfChars - allowedLength)
	} else {
		return getFactorial(nrOfChars)	
	}

}

// CalculatePossibleCombinations returns all the calculated combinations.
func CalculatePossibleCombinations(inputSlice []string, allowedLength int) []string {
	resetGlobalVariables()
	setMaxAllowedStringLength(allowedLength, len(inputSlice))
	var defaultElement = ""
	permute(inputSlice, defaultElement)

	return generatedData
}





