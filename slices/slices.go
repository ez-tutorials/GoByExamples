package main

import "fmt"

func main() {

	// slickes are typed only by the elements they contain.
	// use builtin make() to create a slice of stirngs of length 3.
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	// len() returns the length of the slice
	fmt.Println("len: ", len(s))

	// slices support append(), which returns a slice containing one or more
	// new values.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	// slice can be copied, copy()
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	// slices support slice operation
	l := s[2:5]
	fmt.Println("sl2: ", l)

	l = s[:5]
	fmt.Println("sl3: ", l)

	// delcare and initialized in one line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t)

	// slices can be composed into multi-dimensional data structures
	// the length of the inner slices can vary
	// slice vs array:
	//		slice is dynamic
	//		multi-dimensional slices don't have the same sized inner slices
	//		both rendered the same by fmt.Println()
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
