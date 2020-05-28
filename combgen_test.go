package combgen_test

import (
	"combgen"
	"testing"
)

// equal tells whether a and b contain the same elements.
func equal(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

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

	if !equal(correctOutput, result) {
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

	if !equal(correctOutput, result) {
		t.Errorf("TestCombinationGeneratorSymbols failed.")
	}
}

func TestCombCalcRepeat(t *testing.T) {
	inputData := []string{"A", "B", "C"}
	correctOutput := []string{"ABC", "ABA", "ABB", "ACB", "ACA", "ACC", "AAB", "AAC", "AAA",
		"BAC", "BAA", "BAB", "BCA", "BCB", "BCC", "BBA", "BBC", "BBB",
		"CAB", "CAA", "CAC", "CBA", "CBB", "CBC", "CCA", "CCB", "CCC"}
	result := combgen.CalculateCombinations(inputData, 3, true)

	if !equal(correctOutput, result) {
		t.Errorf("TestCombCalcRepeat failed.")
	}
}

func TestCombCalcRepeatDuplicateChar(t *testing.T) {
	inputData := []string{"A", "B", "C", "A"}
	correctOutput := []string{"ABC", "ABA", "ABB", "ACB", "ACA", "ACC", "AAB", "AAC", "AAA",
		"BAC", "BAA", "BAB", "BCA", "BCB", "BCC", "BBA", "BBC", "BBB",
		"CAB", "CAA", "CAC", "CBA", "CBB", "CBC", "CCA", "CCB", "CCC"}
	result := combgen.CalculateCombinations(inputData, 3, true)

	if !equal(correctOutput, result) {
		t.Errorf("TestCombCalcRepeatDuplicateChar failed.")
	}
}
