package combgen

// Global variable for storing the generated data.
var generatedData []string

// Global variable for storing the output string max length.
var outputMaxLength int

// CalculateCombinationsCount return the possible nr of combinations.
func CalculateCombinationsCount(inputSlice[]string, allowedLength int) int {
	resetGlobalVariables()
	nrOfChars := len(inputSlice)
	setMaxAllowedStringLength(allowedLength, len(inputSlice))	

	return getFactorial(nrOfChars) / getFactorial(nrOfChars - outputMaxLength)
}

// CalculateCombinations returns all the calculated combinations.
func CalculateCombinations(inputSlice []string, allowedLength int) []string {
	resetGlobalVariables()
	setMaxAllowedStringLength(allowedLength, len(inputSlice))
	var defaultElement = ""
	permute(inputSlice, defaultElement)

	return generatedData
}





