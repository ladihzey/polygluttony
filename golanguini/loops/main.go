package main

import "fmt"

func main() {
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}
	fmt.Println()

	for i := range colors {
		fmt.Println(colors[i])
	}
	fmt.Println()

	for _, color := range colors {
		fmt.Println(color)
	}
	fmt.Println()

	value := 1
	for value < 10 {
		fmt.Println(value)
		value++
	}
}
