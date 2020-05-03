package combgen_test

import(
	"testing"
	"combgen"
	"fmt"
)

func TestNrOfPossibleCalculations(t *testing.T) {
	inputData := []string{"A", "B", "C", "D"}
	correctOutput := 24
	result := combgen.NrOfPossibleCalculations(inputData)

	if result != correctOutput {
		t.Errorf("Nr of possible calculations test failed. Expected: %d, got: %d.", correctOutput , result)
	}
}

func TestCalculatePossibleCombinations(t * testing.T) {
	inputData := []string{"A", "B", "C", "D"}
	correctOutput := []string{
		"ABCD", "ABDC", "ACBD", "ACDB", "ADBC", "ADCB",
		"BACD", "BADC", "BCAD", "BCDA", "BDAC", "BDCA",
		"CABD", "CADB", "CBAD", "CBDA", "CDAB", "CDBA",
		"DABC", "DACB", "DBAC", "DBCA", "DCAB", "DCBA"}
	result := combgen.CalculatePossibleCombinations(inputData, 0)

	if !combgen.Equal(correctOutput, result) {
		t.Errorf("Calculating possible combinations test failed.")
	}
}

func TestCombinationGeneratorSymbols(t * testing.T) {
	inputData := []string{"!", "%", "&", "*", "/"}
	correctOutput := []string{
		"!%", "!&", "!*", "!/",
		"%!", "%&", "%*", "%/",
		"&!", "&%", "&*", "&/",
		"*!", "*%", "*&", "*/",
		"/!", "/%", "/&", "/*"}
	result := combgen.CalculatePossibleCombinations(inputData, 2)

	if !combgen.Equal(correctOutput, result) {
		t.Errorf("Calculating possible combinations with symbols test failed.")
	}
}

func TestCalculatePossibleCombinationss(t * testing.T) {
	inputData := []string{"A", "B", "C", "D"}
	correctOutput := []string{
		"ABCD", "ABDC", "ACBD", "ACDB", "ADBC", "ADCB",
		"BACD", "BADC", "BCAD", "BCDA", "BDAC", "BDCA",
		"CABD", "CADB", "CBAD", "CBDA", "CDAB", "CDBA",
		"DABC", "DACB", "DBAC", "DBCA", "DCAB", "DCBA"}
	result := combgen.CalculatePossibleCombinations(inputData, 3)

	fmt.Println(result)

	if !combgen.Equal(correctOutput, result) {
		t.Errorf("Calculating possible combinations test failed.")
	}
}





