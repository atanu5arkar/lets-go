package main

import "fmt"

func main() {
	// Literals are untyped?
	
	fmt.Println("Integer:", 1_300_229)
	fmt.Println("Rune:", 'A')

	fmt.Println("Hello,\n\tWorld!")
	fmt.Println(`Hello,
	World!`)

	name := "Pete"
	
}

/*
 A Go module contains source code and an exact specification (go.mod) of its dependencies. The code within a module is organized
 into one or more packages. The main package is where a Go program starts execution.

 Like C, Go requires a semi-colon at the end of every statement. But we should not include them in the source, as the compiler follows
 a set of predefined rules to insert semi-colons automatically.

 In Go, undefined variables are assigned a "zero value" of the appropriate type.
 
 A literal is an explicitly specified number, character, or string.
*/
