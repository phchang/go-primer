package main

import "fmt"

func main() {

	/* GOROOT and GOPATH
	- `go env`
	- GOROOT -> core Go library and compiler
	- GOPATH -> local libraries / binaries / cache live
	*/

	// print functions
	fmt.Println("Hello world")
	fmt.Printf("A formatted string <%v>\n", "here") // format verbs

	/*
		types in Go -> borrowed from C, but simplified
		  - bool
		  - string
		  - int,  int8,  int16,  int32,  int64 -> typically just use int
		  - uint, uint8, uint16, uint32, uint64, uintptr
		  - byte -> alias for uint8
		  - rune -> alias for int32 (Unicode code point)
		  - float32, float64
		  - complex64, complex128

		zero value is empty string
		  - bool -> false
		  - numeric -> 0
		  - string -> ""
		  - pointers -> nil
	*/

	// declaring variables -> var and const
	var foo string = "foo" // same as `var foo = "foo"`
	bar := "bar"
	const pi = 3.14

	fmt.Printf("foo type is %T\n", foo)
	fmt.Printf("bar type is %T\n", bar)
	fmt.Println("pi is approximately ", pi)

	// explicit type conversion
	var x = 5.7 // float64
	var y = int(x)
	fmt.Println(y) // prints "5"
}
