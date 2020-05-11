# Combination generator
Simple package for generating all possible combinations from slice.

## Functions
Both functions accept 2 parameter(slice - string and the max length of the output strings - int)
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


## TODO
1. Rewrite the code using structs rather global variables and functions.
2. Edit README file.


