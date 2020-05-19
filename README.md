# Combination generator
Simple package for generating all possible combinations from slice.

## Functions
Both functions accept 3 parameter: 
	* []string{} - data for the calculation
	* int - calculated combinations digits length(pass 0 for default)
	* bool - if single character can repeat in the combination

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
	combinationCount := combgen.CalculateCombinationsCount(data, 0, false)

	// Returns: []string{"ABC", "ACB", "BAC", "BCA", "CAB", "CBA"}
	combinationList := combgen.CalculateCombinations(data, 0, false)
}
```


