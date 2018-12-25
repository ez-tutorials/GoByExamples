package main

import "fmt"

func main()  {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// if without an else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// Any varialbes declared in this statement are available in all 
	// branches.
	// Note that you don’t need parentheses around conditions in Go, 
	// but that the braces are required.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}
