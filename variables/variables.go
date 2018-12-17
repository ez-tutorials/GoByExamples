package main

import "fmt"

//In `Go`, variables are explicitly declared and used by the compiler.
//`var` declares 1 or more variables.
//`Go` will infer the type of initialized variables.
//Variables value is default to `zero`.
//`:=` syntax is shorthand for declaring and initializing a variable.

func main() {
	// var delares 1 or more variables
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding initialization are zero-valued.
	var e int
	fmt.Println(e)

	// The := syntax is shorthand for declaring and initializing a variable
	f := "short"
	fmt.Println(f)
}
