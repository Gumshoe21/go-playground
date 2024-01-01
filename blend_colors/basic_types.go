package main

import "fmt"

type Color struct {
	name string
}

type Join interface {
	Mix(o Color)
}

type Combine interface {
	Mix(o Color)
}
func (c *Color) Mix(otherColor Color) {
	fmt.Printf("%v mixed with %v = %v-%v!", c, otherColor, c, otherColor)
}

func main() {
	var blue Combine = &Color{"blue"}
	// var green Join = &Color{"green"}
	// green.Mix(*blue.(*Color)) // a type assertion - we access the underlyingi concrete value of type *Color stored in an interface value

	// blueColor, ok := blue.(*Color)
	// if !ok {
	// 	fmt.Println("blue is not a *Color")
	// 	return
	// }

	switch v := blue.(type) {
	case Combine: 
		fmt.Printf("%v is a Combine.", v)
	case Join: 
		fmt.Printf("%v is a Join.", v)
	}
	// Now you can access the name field
	// fmt.Println(blueColor.name)

	// green.Mix(*blueColor)
}