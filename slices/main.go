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

	// Capacity of a slice refers to the number of consecutive memory locations
	// reserved for its elements. It can be greater than its length.
	// When we try to append to a slice that is full, the runtime allocates a new
	// backing array of bigger capacity. All the elements of the slice are copied
	// to the new array upon which thereafter append is applied.
	var seq []int
	seq = append(seq, 2)
	fmt.Println(len(seq), cap(seq))

	// Define a slice with an initial length. All the elements are initialized to the
	// zero-value of the given type.
	var m = make([]int, 8)
	fmt.Println(m)

	// Reset all the elements of a slice to the zero-value of its type.
	fmt.Println(x)
	clear(x)
	fmt.Println(x)

	// Create slice from a slice using Slice Expressions.
	v := []string{"a", "t", "a", "n", "u"}
	v_prime := v[:3]
	fmt.Println("v_prime:", v_prime)
}
