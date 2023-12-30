package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1 << 64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// Zero values
// Variables declared without an explicit initial value are given their zero value.
//The zero value is:
//0 for numeric types,
//false for the boolean type, and
//"" (the empty string) for strings.
func zeroValues() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n",i, f, b ,s)
	// 0 0 false ""
}