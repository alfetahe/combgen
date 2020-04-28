package combgen

// Global variable for storing the generated data.
var generatedData []string

// NrOfPossibleCalculations return the possible nr of combinations.
func NrOfPossibleCalculations(inputSlice[]string) int {
	nrOfChars := len(inputSlice)
	result := 1
	for i := 1; i < nrOfChars+1; i++ {
		result = i * result
	}

	return result
}

// CalculatePossibleCombinations returns all the calculated combinations.
func CalculatePossibleCombinations(input []string) []string {
	var defaultElement = ""
	permute(input, defaultElement)

	return generatedData
}

// Recursive function for looping thru the slice.
func permute(slice []string, lockedElement string) {
	if len(slice) == 2 {
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
	secondPart := lockedElement + slice[0] + slice[1]
	generatedData = append(generatedData, firstPart, secondPart)
}


