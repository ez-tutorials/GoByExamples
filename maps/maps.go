package main

// maps or hashes or dicts

import "fmt"

func main() {

	// use builtin make() to create empty map
	// make(map[key-type]value-type)
	m := make(map[string]int)

	// set key/value pairs
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map :", m)

	// get value for a key
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// len() returns the number of key/value pairs
	fmt.Println("len: ", len(m))

	// delete() removes key/value pairs from a map
	delete(m, "k2")
	fmt.Println("map: ", m)

	// the optional 2nd return value when getting a value from a map indicates
	// if the key was present in the map.
	// This can be used to disambiguate between missing keys and keys with zero
	// values like 0 or "".
	// _, blank identifier
	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	// declare and initialize in one line
	n := map[string]int{"foo": 1, "bar": 2}
	// maps appear in the form: map[k:v, k:v] when printed with fmt.Println()
}
