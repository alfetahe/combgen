package combgen_test

import(
	"testing"
	"combgen"
	"combgen/helpers"
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
		"ABCD", "ABCD", "ACBD", "ACBD", "ADBC", "ADBC",
		"BACD", "BACD", "BCAD", "BCAD", "BDAC", "BDAC",
		"CABD", "CABD", "CBAD", "CBAD", "CDAB", "CDAB",
		"DABC", "DABC", "DBAC", "DBAC", "DCAB", "DCAB"}
	result := combgen.CalculatePossibleCombinations(inputData)

	if !helpers.Equal(correctOutput, result) {
		t.Errorf("Calculating possible combinations test failed.")
	}
}



