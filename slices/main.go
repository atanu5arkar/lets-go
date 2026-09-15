package main

import "fmt"
import "slices"

func main() {
	// A slice literal is quite similar to an array, except only that
	// we don't specify the size here. Slices grow on-demand!
	var x = []int{3, 4, -1}
	var y = []int{3: 3, 6: 2}

	fmt.Println(x)
	fmt.Println(y)

	// The zero-value for a slice is nil. It is an "untyped" identifier that
	// represents the absence of a value for some types.
	var z []int
	fmt.Println(z)

	// Logical operators are not allowed to test the equality of two slices.
	fmt.Println(z == nil)
	fmt.Println(slices.Equal(x, y))

	// len works for both slice and array types.
	fmt.Println("Size of y:", len(y))

	// The first argument in append creates a copy of the source slice to which 
	// the subsequent elements are appended. This new slice is then returned. 
	// It is mandatory to assign the return value.
	z = append(z, 2, 3)
	x = append(x, y...)

	fmt.Println(z)
	fmt.Println(x)
}
