// Go communicates errors via an explicit, separate return value
// it's easy to see which functions return errors and to handle them using the
// same language constructs employed for any other, non-error tasks.

package main

import "errors"
import "fmt"

// errors are the last return value and have type error, a built-in interface
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New() constructs a basic error value with the given message
		return -1, errors.New("can't work with 42")
	}
	// A nil value in the error position indicates that there was no error
	return arg + 3, nil
}

/*
customized type can be used as erros by implementing the Error() method on them.
Here is a variant on the example above that uses a custom type to explicityly represent an augument error.
*/
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// in case &argError syntax is used to build a new struct, supplying
		// values for the two fields args and prob
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	// 2 loops below test out each of our error-returning funcitons.
	// the use of an inline error check on the if lne is a common idiom in Go
	// code.
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// if you want to programmatically use the data in a custom error, you 'll
	// need to get the error as an instane of the custom error type via type
	// assertion.
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
