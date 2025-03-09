package main

import (
	"fmt"
	"strings"
)

const totalIterations = 5

func main() {
	fmt.Println(Repeat("a"))
}

// notes strings are immutable
// meaning every concatenation, such as in our Repeat function,
// involves copying memory to accommodate the new string.
// This impacts performance, particularly during heavy string concatenation.
// dan setelah dicoba, ketika menggunakan string dan strings.Builder perbedaannya 137.7 ns vs 33.10 ns
func Repeat(character string) string {
	//var characters string
	var repeated strings.Builder
	for i := 0; i < totalIterations; i++ {
		//	characters += character
		repeated.WriteString(character)
	}
	//return characters
	return repeated.String()
}
