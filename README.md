# Combination generator
Simple package for generating all possible combinations from slice.

## Functions
Both functions accept 2 parameter a slice of strings and int value of how many characters the output will be(pass 0 for default).
1. CalculateCombinationsCount - returns integer value of possible combinations count.
2. CalculateCombinations - returns a slice of string which contains all possible combinations.

# Installation
``` bash
go get github.com/anuarsaeed/combgen
```
Command above will download the package to your $GOPATH/src/github.com/ directory.


# Usage example
``` go
package main

import(
    "github.com/anuarsaeed/combgen"
)

func main() {
	data := []string{"A", "B", "C"}

	// Returns: 6
	combinationCount := combgen.CalculateCombinationsCount(data, 0)

	// Returns: []string{"ABC", "ACB", "BAC", "BCA", "CAB", "CBA"}
	combinationList := combgen.CalculateCombinations(data, 0)
}
```


