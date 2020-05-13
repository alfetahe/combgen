package combgen

type combgenInterface interface {
	calculateNrOfCombinations() int
	calculateCombinations() []string
}