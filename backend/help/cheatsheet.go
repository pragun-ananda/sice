/*
	General:
		- Modules: When your code imports packages from other modules, you manage that in your own module
			- usually your module is created where your github repository is kept github.com/mymodule (initiate with "go mod init github.com/mymodule")
			- when importing an external package use pkg.go.dev and run "go mod tidy" to add package requirements to go.mod

		- A tour of Go: https://tour.golang.org/

*/

// Single-line comments are written with //

/*
	Multi-line comments are written like this (same as Java/JavaScript)
*/

// package main causes the main method to run automatically
package main

import "fmt"

/*
Can have multiple imports using import (
	"fmt"
	t "time"	<- imports time with the package alias t (can have aliases for variables too)
)
*/

/*
	Function Definitions:
		- func methodName(varName inputType) outputType {

		}
		- Sprintf returns a string, plugging in variables in place of format specifiers
		- Printf prints the same string
*/
func hello(name string) string {
	message := fmt.Sprintf("Hi! %v. Welcome!", name)
	return message
}

// func defines methods
func main() {
	fmt.Println(hello("Bob"))

	// define variables with the var keyword (ex: var variableName dataType)
	// data types include string, int (can specify 8-byte w/ int8), bool, float32 (or float64)
	var helloWorld string = "test"

	// can also define variables w/ type inference using the walrus operator
	y := 3

	// can have multiple variable declarations too
	var a, b, c = 1, 2, 4

	helloWorld = "Hello World"

	fmt.Println(helloWorld)
	fmt.Println(y)
	fmt.Println(a, b, c)

	fmt.Println(quote.Go())
}

/*
Go Basic Types

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     represents a Unicode code point

float32 float64

complex64 complex128

Unitialized variables are given zero types:
- 0 for numeric types,
- false for the boolean type, and
- "" (the empty string) for strings.

Can convert types: var z uint = uint(f) or z := uint(f)

Constants - use const keyword: const Pi = 3.14 (cannot use walrus operator)
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

There's only one way to loop in Go - for loops

for i := 0; i < 10; i++ {
		sum += i
}

structure:
for (init); (condition); (post) {
	// the init is optional
	// the post is optional
}

// While loop - can drop the semicolons
for sum < 1000 {
		sum += sum
}

// infinite loop
for {
}

// If Statements need to have braces
if x < 0 {
		return sqrt(-x) + "i"
}

// If statements can also have init conditions
if v := math.Pow(x, n); v < lim {
		return v
}
return lim

// Switch statemetns
switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
}

// Deferred functions are pushed onto a stack for when the function returns
for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

Go Pointers work the same as way as in C/C++:
p := &i         // point to i
fmt.Println(*p) // read i through the pointer
*p = 21         // set i through the pointer
fmt.Println(i)  // see the new value of i

p = &j         // point to j
*p = *p / 37   // divide j through the pointer
fmt.Println(j) // see the new value of j

A struct is a collection of fields:
type Vertex struct {
	X int
	Y int
}

You can access structs through pointers (automatic dereference)
v := Vertex{1, 2}
p := &v
p.X = 1e9

Defining arrays:
var a [10]int
primes := [6]int{2, 3, 5, 7, 11, 13}

Indexing is like Python
primes := [6]int{2, 3, 5, 7, 11, 13}
var s []int = primes[1:4]

Slices do not store copies of an array. They store references
Changing a slice, changes the underlying array

Writing Web Applications: https://golang.org/doc/articles/wiki/
*/
