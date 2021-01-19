package main

import "fmt"

// just declaring a variable  with the package scope
var message = `Hello there ! What are you thinking about Go ?
It is a proper language for you ?
`
// Notice that I used `. When you use that, the string collected is kept as passed
// to the variable.

func main() {
	fmt.Print(message)
}