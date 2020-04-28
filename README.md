# Combination generator
Simple package for generating all possible combinations from slice.

## Functions
Both functions accept 1 parameter(slice of strings)
1. NrOfPossibleCalculations - returns integer value of possible combinations.
2. CalculatePossibleCombinations - returns a slice of string which contains all possible combinations.

# Installation
go get github.com/anuarsaeed/combgen
Command above will download the package to your $GOPATH/src/github.com/ directory.

# Usage example
package main

import(
    "github.com/anuarsaeed/combgen"
	"fmt"
)

func main() {
	data := []string{"A", "B", "C"}

	// Outputs: 6
	fmt.Println(combgen.NrOfPossibleCalculations(data))

	// Outputs: [ABC ABC BAC BAC CAB CAB]
	fmt.Println(combgen.CalculatePossibleCombinations(data))
}

