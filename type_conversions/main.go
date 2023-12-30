package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	// var z uint = f  <-- "cannot use f (variable of type float64) as type uint in variable declaration"
	fmt.Println(x, y, z)
}