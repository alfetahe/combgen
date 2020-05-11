# Combination generator
Simple package for generating all possible combinations from slice.

## Functions
Both functions accept 2 parameter a slice of strings and int value of how many characters will be used.
1. NrOfPossibleCalculations - returns integer value of possible combinations.
2. CalculatePossibleCombinations - returns a slice of string which contains all possible combinations.

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
	"fmt"
)

func main() {
	data := []string{"A", "B", "C"}

	// Outputs: 6
	fmt.Println(combgen.NrOfPossibleCalculations(data, 0))

	// Outputs: [ABC ACB BAC BCA CAB CBA]
	fmt.Println(combgen.CalculatePossibleCombinations(data, 3))
}
```


