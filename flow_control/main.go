package main

import (
	"fmt"
	"math"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum) // 45
}

// The init and post statements are optional
// Essentially a "while" loop in other langs
func optional() {
	sum := 1

	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

// For is Go's "while"
// At that point you can drop the semicolons: C's while is spelled for in Go.

func gosWhile() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// Forever
// Omitting the loop cond. makes it loop forever (infinite loop)

func forever() {
	for {

	} // timeout running program
}

// If
// Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func sqrtMaini() {
	fmt.Println(sqrt(2), sqrt(-4))
}

// If with a short := statement

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
	return v // this would not work as v is out of scope as soon as the if statement terminates
}

func pow_without_short(x, n,lim float64) float 64 {
	v := math.Pow(x, n)
	if v < lim {
		return v
	}
	return lim
}

// If and else

func pow_else(x,n,lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %\n", v, lim)
	}
	// can't use v here
	return lim
}

// Switch
// Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

func switch_statements() {
	fmt.Printi("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

// Switch eval order

// Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

func switch_eval_order() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// Switch without a condition is the same as switch true.

// This construct can be a clean way to write long if-then-else chains.
func switch_with_no_cond() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// Defer

func defer_statements() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

// Stacking Defers
// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
func stacking_defers() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}