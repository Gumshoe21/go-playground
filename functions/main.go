package main

import "fmt"

func add(x int, y int) int {
	return x + y
}


func add_with_omitted_type(x, y int) int {
	return x + y
}

func swap(x string, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(42, 13))
	y, x := swap("hello", "world")
	fmt.Println(y,x)
}

