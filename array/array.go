package main

import (
	"fmt"
)

func main() {
	// create array holds exactly 5 ints
	// the type of elements and lenght are both part of the array's type
	var a [5]int
	// by defaut, an array is zero-valued
	fmt.Println("emp: ", a)

	// set/get a value at an index
	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	// the builtin len() returns the length of an array
	fmt.Println("len: ", len(a))

	// declare and initialize an array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	// arrays are 1-dimensional, but multi-dimentional array can be built
	// with composed types
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
