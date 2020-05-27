package combgen

type combgenInterface interface {
	calculateNrOfCombinations() int
	calculateCombinations() []string
	permute(string, []string)
}
