package main

import (
	"fmt"
	"math"
)

func main() {
	// go does not implicitly convert types
	// e.g. float64 must be explicitly converted to int and vice versa
	integer := 10
	float := 10.9

	intSum := integer + int(float)
	fmt.Printf("Sum: %v, Type: %T\n", intSum, intSum)

	floatSum := float + float64(integer)
	fmt.Printf("Sum: %v, Type: %T\n", floatSum, floatSum)
	fmt.Println()

	f1, f2, f3 := 10.5, 20.1, 30.3
	sum := f1 + f2 + f3
	fmt.Printf("Sum: %v, Type: %T\n", sum, sum)

	sum = math.Round(sum * 100) / 100
	fmt.Printf("Sum: %v, Type: %T\n", sum, sum)
	fmt.Println()

	circleRadius := 15.5
	circumference := 2 * math.Pi * circleRadius
	fmt.Printf("Circumference: %.2f\n", circumference)
	fmt.Println()
}
