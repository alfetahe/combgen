package combgen_test

import (
	"combgen"
	"testing"
)

func TestCombCount(t *testing.T) {
	inputData := []string{"A", "B", "C", "D"}
	correctOutput := 24
	result := combgen.CalculateCombinationsCount(inputData, 4, false)

	if result != correctOutput {
		t.Errorf("TestCombCount failed. Expected: %d, got: %d.", correctOutput, result)
	}
}

func TestCombCountDefaultLength(t *testing.T) {
	inputData := []string{"A", "B", "C", "D"}
	correctOutput := 24
	result := combgen.CalculateCombinationsCount(inputData, 0, false)

	if result != correctOutput {
		t.Errorf("TestCombCountDefaultLength failed. Expected: %d, got: %d.", correctOutput, result)
	}
}

func TestCombCountCharRepeat(t *testing.T) {
	inputData := []string{"A", "B", "C", "D"}
	correctOutput := 1024
	result := combgen.CalculateCombinationsCount(inputData, 5, true)

	if result != correctOutput {
		t.Errorf("TestCombCountCharRepeat failed. Expected: %d, got: %d.", correctOutput, result)
	}
}

func TestCombCountCharRepeatDefaultLength(t *testing.T) {
	inputData := []string{"A", "B", "C", "D"}
	correctOutput := 256
	result := combgen.CalculateCombinationsCount(inputData, 0, true)

	if result != correctOutput {
		t.Errorf("TestCombCountCharRepeatDefaultLength failed. Expected: %d, got: %d.", correctOutput, result)
	}
}

func TestCombCalc(t *testing.T) {
	inputData := []string{"A", "B", "C", "D"}
	correctOutput := []string{
		"ABCD", "ABDC", "ACBD", "ACDB", "ADBC", "ADCB",
		"BACD", "BADC", "BCAD", "BCDA", "BDAC", "BDCA",
		"CABD", "CADB", "CBAD", "CBDA", "CDAB", "CDBA",
		"DABC", "DACB", "DBAC", "DBCA", "DCAB", "DCBA"}
	result := combgen.CalculateCombinations(inputData, 0, false)

	if !combgen.Equal(correctOutput, result) {
		t.Errorf("TestCombCalc failed.")
	}
}

func TestCombCalcSymbols(t *testing.T) {
	inputData := []string{"!", "%", "&", "*", "/"}
	correctOutput := []string{
		"!%", "!&", "!*", "!/",
		"%!", "%&", "%*", "%/",
		"&!", "&%", "&*", "&/",
		"*!", "*%", "*&", "*/",
		"/!", "/%", "/&", "/*"}
	result := combgen.CalculateCombinations(inputData, 2, false)

	if !combgen.Equal(correctOutput, result) {
		t.Errorf("TestCombinationGeneratorSymbols failed.")
	}
}
