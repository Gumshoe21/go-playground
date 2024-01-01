package main

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"time"
)

type Vertex struct {
	X, Y float64
}

// Receiver methods

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func methods_main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

// Non-struct type methods

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).

func nonStruct() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

// Pointer receivers
// functions with a pointer argument must take a pointer:
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func pointerReceivers() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

// Methods and pointer indirection

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func pointerIndirection() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	// p := Vertex{4, 3} <-- cannot use p (variable of type Vertex) as *Vertex value in argument to ScaleFunc (i.e. we can't use Vertex as a value in a function without receivers)
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

// Interfaces

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func interfaceMain() {
	var i I = T{"hello"}
	i.M()
}


type myInterface interface {
	printHelloWorld()
}

type myStruct struct {
	S string
	I int
}

func (m myStruct) printHelloWorld() {
	fmt.Printf("%v - %v", m.I, m.S)
}

func myInterfaceFunc() {
	var coolInterface myInterface = myStruct{"Sup", 1}
	coolInterface.printHelloWorld()
}


// Interface values with nil underlying values

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func underlying() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// Nil interface values

// Empty interface

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// (<nil>, <nil>)
// (42, int)
// (hello, string)

// Type Assertions
func typeAssertions() {

var i interface{} = "hello"

s := i.(string)

s, ok := i.(string)

f, ok := i.(float64)

f = i.(float64)
}

// Type switches

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// Errors

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError {
		time.Now(),
		"it didn't work",
	}
}

func errorsMain() {
	if err := run; err != nil {
	fmt.Println(err)
}
}
func errorsBasics() {
	i, err = strconv.Atoi("42")
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
	}
	fmt.Println("Converted integer:", i)
}

// Readers

func readersBasics() {
	r := strings.NewReader("Hello, Reader!")
	b:= make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}