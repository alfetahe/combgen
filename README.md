# Combination generator
Simple package for generating all possible combinations from slice.

## Functions
1. CalculateCombinationsCount([]string{}, int, bool) int - returns integer value of possible combinations count.
2. CalculateCombinations([]string{}, int, bool) []string - returns a slice of string which contains all possible combinations.

Both functions accept 3 parameter: 
1. []string{} - data for the calculation
2. int - calculated combinations character length(pass 0 for default, will be the length of input slice)
3. bool - if character can repeat

# Installation
``` bash
go get github.com/alfetahe/combgen
```
Command above will download the package to your $GOPATH/src/github.com/ directory.


# Usage example
``` go
package main

import(
    "github.com/alfetahe/combgen"
)

func main() {
	data := []string{"A", "B", "C"}

	// Returns: 6
	combgen.CalculateCombinationsCount(data, 0, false)

	// Returns: []string{"ABC","ACB","BAC","BCA","CAB","CBA"}
	combgen.CalculateCombinations(data, 0, false)

	// Returns: 27
	combgen.CalculateCombinationsCount(data, 0, true)

	// Returns: []string{"ABC","ABA","ABB","ACB","ACA","ACC","AAB","AAC","AAA","BAC","BAA","BAB","BCA","BCB","BCC","BBA","BBC","BBB","CAB","CAA","CAC","CBA","CBB","CBC","CCA","CCB","CCC"}
	combgen.CalculateCombinations(data, 0, true)
}
```


