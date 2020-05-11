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

func getFactorial(nrOfChars int) int {
	factorial := 1;
	for i := 1; i <= nrOfChars; i++ {
		factorial = i * factorial
	}

	return factorial
}

// CalculatePossibleCombinations returns all the calculated combinations.
func CalculatePossibleCombinations(input []string, allowedLength int) []string {
	resetGlobalVariables()
	setMaxAllowedStringLength(allowedLength, len(input))
	var defaultElement = ""
	permute(input, defaultElement)

	return generatedData
}

// Recursive function for looping thru the slice.
func permute(slice []string, lockedElement string) {
	if outputMaxLength == len(lockedElement) {
		generatedData = append(generatedData, lockedElement)
	} else if len(slice) == 2 && len(lockedElement) > outputMaxLength {
		twoElSliceWorker(slice, lockedElement)
	} else {
		for index, lockedEl := range slice {
			newSlice := RemoveFromSliceByIndex(slice, index)
			permute(newSlice, lockedElement+lockedEl)
		}
	}
}

// Function for populating the generated data global variable.
func twoElSliceWorker(slice []string, lockedElement string) {
	firstPart := lockedElement + slice[0] + slice[1]
	secondPart := lockedElement + slice[1] + slice[0]
	generatedData = append(generatedData, firstPart, secondPart)
}

// Set max allowed string length.
func setMaxAllowedStringLength(allowedLength int, inputLength int) {
	result := checkIfAllowedStringLengthIsValid(allowedLength, inputLength)
	if result == true {
		outputMaxLength = allowedLength
	} else {
		outputMaxLength = inputLength
	}
}

func checkIfAllowedStringLengthIsValid(allowedLength int, inputLength int) bool {
	if allowedLength == 0 || allowedLength > inputLength {
		return false
	} else {
		return true
	}
}

func resetGlobalVariables() {
	generatedData = []string{}
	outputMaxLength = 0
}

