package main

import (
	"fmt"
)

func main() {
	// When the right hand side of the declaration is typed, the new var is of that same type:
	v := 42 // change me!
	i := v
	fmt.Printf("v is of type %T\n", v) // v is of type int
	fmt.Printf("i is of type %T\n", i) // i is of type int

	// But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant:

	j := 42 // int
	f := 3.142 // float64
	g := 0.867 + 0.5i // complex128
}


