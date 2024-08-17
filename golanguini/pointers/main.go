package main

import "fmt"

func main() {
	i := 42
	p1 := &i
	fmt.Println("Integer pointer:", *p1)

	f := 42.13
	p2 := &f
	fmt.Println("Float pointer:", *p2)

	*p1 = *p1 / 31
	fmt.Println("Value:", i)
	fmt.Println("Pointer:", *p1)
}
