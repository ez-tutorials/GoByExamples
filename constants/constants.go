package main

import (
	"fmt"
	"math"
)

// `const` declares a constant value.
// It can appear anywhere in a `var` statement can.
const s string = "constant"

func main()  {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	// A numeric constant has no type until it's given one by an explicit cast.
	fmt.Println(int64(d))

	// A number can be given a type by using it in a context that requires one.
	// Here n is casted to float64 explicitly
	fmt.Println(math.Sin(n))
}