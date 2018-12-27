// Go's structs are typed collections of fields.
// they are useful for grouping data togeter to form records.

package main

import "fmt"

// this person struct has name and age
type person struct {
	name string
	age  int
}

func main() {
	// create new struct instance
	fmt.Println(person{"Bob", 20})
	// name the fields during initialization
	fmt.Println(person{name: "Alice", age: 30})
	//omitted fields will be zero-valued
	fmt.Println(person{name: "Fred"})

	fmt.Println(person{name: "Ann", age: 50})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	// pointer to a struct
	// access fields with dot
	sp := &s
	fmt.Println(sp.age)

	// struct pointer automatically dereferenced.
	// struct is mutable1
	sp.age = 51
	fmt.Println(s.age)
}
