package main

import "fmt"

func main() {
	counter := counter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
