package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)	// 43
	fmt.Println(&a) // mem. addr.

	var b = &a
	fmt.Println(b)	// mem. addr.
	fmt.Println(*b)	// 43

	// the asterisk dereferences, that is, shows the stored value
	// at stored memory address.
}
