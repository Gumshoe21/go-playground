package main

import "fmt"

// A var statement can be at package or function level.
var c, python, java bool // type is last

func main() {
	var i int
	fmt.Println(i, c, python, java)
	nilTest()
}

// Variables with initializers

var i, j int = 1, 2

func varsWithInits() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}


// Short var decs (short assignment statement :=)
// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

func shortVarDecs() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i,j,k,c,python,java)
}

func nilTest() {
	nil := 123 
	fmt.Println(nil)
}

