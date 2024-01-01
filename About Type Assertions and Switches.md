During my journey in learning about Go, I've come across a confusing yet interesting quirk regarding type assertions.

I've learned that when running a `switch` statement on a value's type, **you can match a type that is different than the value's declared type if the value's underlying type satisfies the matched type.**

Take the following code for example:

```go
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

	switch v := blue.(type) {
	case Join: 
		fmt.Printf("%v is a Join.", v)
	case Combine: 
		fmt.Printf("%v is a Combine.", v)
	}
}
```

Here, we have a `Color` struct and two interfaces, `Join` and `Combine` that both have the same name method, `Mix` as their specification.

Below these definitions, we define a function `Mix` whose receiver type is `*Color`. This means that `*Color` satisfies both the `Join` and `Combine` interfaces due to the `Mix` method. 

Now, here's where things get confusing (and interesting):

```go
func main() {
	var blue Combine = &Color{"blue"}

	switch v := blue.(type) {
	case Join: 
		fmt.Printf("%v is a Join.", v)
	case Combine: 
		fmt.Printf("%v is a Combine.", v)
	}

}
```

What do you think the output of this code should be? This, right?

```&{blue} is a Combine```

Wrong! We get this instead:

```&{blue} is a Join```

But why? The answer lies in what I stated before:

"**You can match a type that is different than the value's declared type if the value's underlying type satisfies the matched type.**"

So let's go over this in greater detail:

Here are the facts:

- The value whose type we're matching: `blue`
- The value's declared type: `Combine`
- The value's underlying type: `*Color`
- The type the value matched: `Join`

Because of the fact that the value's (`blue`) underlying type (`*Color`) satisfies the matched type (`Join`), it will match it and not the expected type (i.e., the declared type `Combine`). It's really that simple.