# Combination generator
Simple package for generating all possible combinations from slice.

## Functions
Both functions accept 3 parameter: 
### []string{} - data for the calculation
### int - calculated combinations character length(pass 0 for default)
### bool - if character can repeat

1. CalculateCombinationsCount() - returns integer value of possible combinations count.
2. CalculateCombinations() - returns a slice of string which contains all possible combinations.

# Installation
``` bash
go get github.com/alfetahe/combgen
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
	combgen.CalculateCombinationsCount(data, 0, false)

	// Returns: []string{"ABC", "ACB", "BAC", "BCA", "CAB", "CBA"}
	combgen.CalculateCombinations(data, 0, false)

	// Returns: 27
	combgen.CalculateCombinationsCount(data, 0, true)
}
```


