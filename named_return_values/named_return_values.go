package main

import "fmt"

func split(sum int) (x int, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return
}

func maini() {
	fmt.Println(split(17))
}