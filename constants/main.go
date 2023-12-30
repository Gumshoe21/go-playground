package main

import "fmt"

// Constants cannot be declared using the := syntax
const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

// Numeric constants

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100 // 1 + 100 zeroes
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99 // 2 
)

func needInt(x int) int { return x*10 + 1}
func needFloat(x float64) float64 {
	return x * 0.1
}

func numericConstants(){
	fmt.Println(needInt(Small)) // needInt(2)
	fmt.Println(needFloat(Small)) // needFloat(2)
	fmt.Println(needFloat(Big)) // needFloat(1 + 100 zeroes)
}
