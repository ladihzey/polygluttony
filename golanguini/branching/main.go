package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	number := 42
	if number < 0 {
		fmt.Println("Number is negative")
	} else if number == 0 {
		fmt.Println("Number is zero")
	} else {
		fmt.Println("Number is positive")
	}

	// it's possible to declare a variable in the if statement
	// this variable is only available in the scope of the if statement
	if num, err := strconv.ParseInt("a256", 10, 64); err == nil {
		fmt.Println(num)
	} else {
		fmt.Println("Error parsing number")
	}

	dow := rand.Intn(7) + 1
	fmt.Println("Day of week:", dow)

	var dayName string
	switch dow {
	case 1:
		dayName = "Monday"
		// break keyword is not needed in Go
	case 2:
		dayName = "Tuesday"
		fallthrough // works like absence of break in C
	default:
		dayName = "Dunno what that day is..."
	}
	fmt.Println("Day of week:", dayName)
}
