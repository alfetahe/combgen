package combgen_test

import(
	"testing"
	"combgen"
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
	inputData := []string{"A", "B", "C"}
	correctOutput := []string{
		"ABCD", "ABDC", "ACBD", "ACDB", "ADBC", "ADCB",
		"BACD", "BADC", "BCAD", "BCDA", "BDAC", "BDCA",
		"CABD", "CADB", "CBAD", "CBDA", "CDAB", "CDBA",
		"DABC", "DACB", "DBAC", "DBCA", "DCAB", "DCBA"}
	result := combgen.CalculatePossibleCombinations(inputData)

	if !combgen.Equal(correctOutput, result) {
		t.Errorf("Calculating possible combinations test failed.")
	}
}



