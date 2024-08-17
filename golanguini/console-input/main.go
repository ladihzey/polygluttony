package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Hello,", input)

	fmt.Print("Enter your age: ")
	ageInput, _ := reader.ReadString('\n')
	age, err := strconv.Atoi(strings.TrimSpace(ageInput))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Printf("You are %d years old\n", age)
}
