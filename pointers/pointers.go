package main

import "fmt"

// zeroval has an int parameter, so arguments will be passed to it by value
// it will get a copy of ival distinct from the one in the calling function
func zeroval(ival int) {
	ival = 0
}

// zeroptr has an *int parameter, it takes an int pointer.
// the *iptr code nt the function body dereferences the pointer from its memory address to the
// current value at that address.
// assigning a value to a dereferenced pointer changes the value at referenced
// address.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	// zeroptr changes the i because it has a reference to the memory address
	// for
	zeroptr(&i)
	fmt.Println("zeroptr: ", i)

	// the &i syntax gives the memory address of i, i.e. a pointer to i
	fmt.Println("pointer: ", &i)

	fmt.Println("i: ", i)
}
