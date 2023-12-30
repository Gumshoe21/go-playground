package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}


/*
"In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.

pizza and pi do not start with a capital letter, so they are not exported.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

Run the code. Notice the error message.

To fix the error, rename math.pi to math.Pi and try it again."

==============

This is from the Tour of Go from Golang. Here is the provided code:

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}


My questions:

Does this mean that a package called "pi" could potentially exist in the math package but we would not be able to export it because it's not exported? I.e., it's inaccessible because that's what the package authors intended?

**********************
		ANSWER:
**********************

In Go, the visibility of a name (variable, function, etc.) is determined by whether its first letter is uppercase or lowercase. If the first letter is uppercase, the name is exported and accessible from outside the package. If the first letter is lowercase, the name is not exported and is only accessible within the package.

In the provided example, math.Pi is accessible because Pi starts with an uppercase letter, making it an exported name. However, if there were a variable named pi (with a lowercase "p") inside the math package, it would not be accessible from outside the package because it would be unexported.

So, to answer your question: Yes, it is possible that an unexported variable named pi could exist within the math package, but it would not be accessible from outside the package. The package authors may choose to export only certain names to control what is visible and accessible to users of the package.
*/

