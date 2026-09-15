/*
 Go considers the size of an array as part of its type. No type casting is possible
 to make arrays of different sizes have the same type. It prevents having functions 
 that can work with an array argument of any size.
*/

package main

import "fmt"

func main() {
	// An integer array of size 3. By default, its elements are initialized
	// to the zero value of the specified type, 0.
	var x [3]int
	fmt.Println(x)

	// An array literal
	var y = [3]int{4, 1, 9}
	fmt.Println(y)

	// We can have a sparse array by defining non-zero values at specific
	// indexes only.
	var z = [10]int{1: 2, 9: 3}
	fmt.Println(z)

	// Arrays are comparable using the == and != operators. Two are equal if
	// they have the same length and the same values.
	fmt.Println(x == y)

	// Multi-dimensional array
	var m = [3][2]int{1: {3, 7}}

	fmt.Println(m)
	fmt.Println(m[1])
}
