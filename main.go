package main

import (
	"fmt"
	"math/rand"
)

// package name is same as last element of import path - so here it is "package rand"

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
