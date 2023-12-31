package main

import "fmt"

func main() {

}

// Pointers

func pointers_basics() {
	// The type *T is a pointer to a T value. Its zero value is nil.
	// var p *int

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

// Structs

type Vertex struct {
	X int
	Y int
}

func structs_basics() {
	fmt.Println(Vertex{1, 2})
	
	// Struct fields
	v := Vertex{1,2}
	v.X = 4
	fmt.Println(v.X) // 4
}

// Pointers to structs
// Struct fields can be accessed through a struct pointer.

// To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
func pointers_to_structs() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

// Struct literals

func struct_literals() {
	var (
		v1 = Vertex{1, 2} // has type Vertex
		v2 = Vertex{X: 1} // Y:0 is implicit
		v3 = Vertex{} // X:0 and Y:0
		p = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1,p,v2,v3) // {1 2} &{1 2} {1 0} {0 0}
}

// Arrays
// An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.

func arrays_basics() {
	var a [2]string // an array of 2 string values
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2,3,5,7,11,13}
	fmt.Println(primes)
}

// Slices
// An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays. The type []T is a slice with elements of type T. A slice is formed by specifying two indices, a low and high bound, separated by a colon: a[low : high] This selects a half-open range which includes the first element, but excludes the last one. The following expression creates a slice which includes elements 1 through 3 of a: a[1:4]

func slices_basics() {
	primes := [6]int{2,3,5,7,11,13}
	var s []int = primes[1:4]
	fmt.Println(s) // [3 5 7]
}

// Slices are like refs to arrays

func slices_are_like_refs_to_arrays() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// [John Paul George Ringo]
	// [John Paul] [Paul George]
	// [John XXX] [XXX George]
	// [John XXX George Ringo]
/
}

// Slice literal
// A slice literal is like an array literal without the length.
var array_literal = [3]bool{true, true, false}
var slice_literal = []bool{true,true,false}


func slice_literals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	r := []bool{true,false,true,true,false,true}
	fmt.Prinitln(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false}
	}
	fmt.Println(s)
}

// Slice defaults
func slice_defaults() {
	s := []int{2,3,5,7,11,13}

	s = s[1:4]

	s = s[:2]

	s = s[1:]
}

func slice_len_and_cap() {
	s := []int{2,3,5,7,11,13} // [2 3 5 7 11 13]
	s = s[:0] // []
	s = s[:4] // [2 3 5 7]
	s = s[2:] // [5 7]
}

// The zero value of a slice is nil. A nil slice has a length and capacity of 0 and has no underlying array.

func slice_zero_val() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func slices_with_make() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func slices_of_slices() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	// X _ X
	// O _ X
	// _ _ O
}

// Appending to a slice
func appending_slices() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// len=0 cap=0 []
// len=1 cap=1 [0]
// len=2 cap=2 [0 1]
// len=5 cap=6 [0 1 2 3 4]

// Range

func range_basics() {
	var pow = []init{1, 2, 4, 8, 16, 32, 64, 128}

	for index, value := range pow {
		fmt.Printf("2**%d = %d\n", index, value)
	}
	// 2**0 = 1
	// 2**1 = 2
	// 2**2 = 4
	// 2**3 = 8
	// 2**4 = 16
	// 2**5 = 32
	// 2**6 = 64
	// 2**7 = 128
}
// Skipping idx or val in Range

func skipping() {
	pow := make([]int, 10)
	// Skipping index
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	// Skipping value
	for index, _ := range pow {
		pow[i] = 1 << uint(1) // 2**1
	}
	// If you only want the index, you can omit the second variable.
	for index := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
}

// Maps

type AnotherVertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func maps_basics() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74,39967,
	}
	fmt.Println(m["Bell Labs"])
}

func map_literals() {
	var m = map[string]Vertex{
		"Bell Labs": Vertex{1,1},
		"Google": Vertex{1,1},
	}
	fmt.Println(m)
}

// If the top-level type is just a type name, you can omit it from the elements of the literal.

func omit_type() {
	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
}

// Mutating maps

func mutating_maps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	// If key is in m, ok is true. If not, ok is false.
	// If key is not in the map, then elem is the zero value for the map's element type.
	elem, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}

// The value: 42
// The value: 48
// The value: 0
// The value: 0 Present? falseu

// Function values

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func compute_main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5,12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

// Function closures

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func adder_main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Prinitln(
			pos(i),
			neg(-2*i),
		)
	}
}