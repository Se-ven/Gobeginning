package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b = &a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 42
	fmt.Println(a)
}

	// line 16 requests changing the value at said address be changed to 42.
	// Addresses to change/pass information is useful...when?