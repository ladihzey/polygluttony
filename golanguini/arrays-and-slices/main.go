package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)
	fmt.Println(colors[2])
	fmt.Println()

	numbers := [5]int{5, 10, 15, 20, 25}
	fmt.Println(numbers)
	fmt.Println("Number of numbers:", len(numbers))
	fmt.Println()

	colorsSlice := []string{"Red", "Green", "Blue", "Yellow", "Pink"}
	fmt.Println(colorsSlice)
	colorsSlice = append(colorsSlice, "Purple")
	fmt.Println(colorsSlice)

	colorsSlice = append(colorsSlice[1:len(colorsSlice)])
	fmt.Println(colorsSlice)

	colorsSlice = append(colorsSlice[:len(colorsSlice) - 1])
	fmt.Println(colorsSlice)

	sort.Strings(colorsSlice)
	fmt.Println(colorsSlice)
}
