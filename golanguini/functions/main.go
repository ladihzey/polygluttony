package main

import "fmt"

func main() {
	sayHello()
}

func sayHello() {
	println("Hello, world!")
	sum := addValue(1, 2)
	fmt.Println(sum)

	sum, count := addValues(1, 2, 3, 4, 5)
	fmt.Println(sum, count)
}

// There is an alternative syntax for parameters, that has the same type
// func addValue(a, b int) int {
func addValue(a int, b int) int {
	return a + b
}

func addValues(values ...int) (int, int) {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum, len(values)
}
