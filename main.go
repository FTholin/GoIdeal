package main

import (
	"GoIdeal/existence"
	"fmt"
)

// TODO need to debug primitives methods with the code or the algorithm
func main() {
	existence := existence.NewExistence010()

	for i := 0; i < 20; i++ {
		stepTrace := existence.Step()
		fmt.Printf("%d: %s\n", i, stepTrace)
	}

	print(existence)
}