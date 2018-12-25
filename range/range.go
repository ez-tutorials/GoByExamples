package main

import "fmt"

/*
range iterates over elements in variety of data structures.
*/

func main() {
	// sum the numbers in a slice or array
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	// range on arrays and slices provides both the index and value for each
	// entry.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	// range on map iterates key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println("%s -> %s\n", k, v)
	}
	// range can also iterates just the keys of a map
	for k := range kvs {
		fmt.Println("key: ", k)
	}

	// range on strings iterates Unicode code points.
	// the 1st value is the starting byte index of the rune and the 2nd the
	// rune iteself
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
