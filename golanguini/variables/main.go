package main

import "fmt"

func printValue(value interface{}) {
	fmt.Print(value)
	fmt.Printf(" :: %T\n", value)
}

func main() {
	var str string = "Hello, World!"
	printValue(str)

	var number int = 42
	printValue(number)

	// go variables could automatically infer the type
	var inferred = 25.5
	printValue(inferred)

	// variables are mutable and could be reassigned
	inferred = 25.3
	printValue(inferred)

	// go does not support shadowing
	// var inferred = 25 -> error

	// there is an alternative syntax for variables definition
	alternative := "alternative syntax"
	printValue(alternative)

	// constants are defined using the const keyword
	const constant = "constant"
	printValue(constant)
}
