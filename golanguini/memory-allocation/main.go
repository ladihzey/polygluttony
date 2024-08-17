package main

import "fmt"

// The Go runtime is statically linked into application.
// Memory is allocated and deallocated by the runtime.
// Use `make()` or `new()` to initialize maps, slices, and channels.

// `new()` allocates but does not initialize memory.
// Results in zeroed storage but returns a memory address.

// `make()` initializes memory and returns a reference to it.
// Allocates non-zeroed storage and returns a memory address.
func main() {
	m := make(map[string]int) // does not work with `new`
	m["one"] = 1
	fmt.Println(m)

	i := 42
	p1 := &i
	fmt.Println("Integer pointer:", *p1)

	f := 42.13
	p2 := &f
	fmt.Println("Float pointer:", *p2)

	*p1 = *p1 / 31
	fmt.Println("Value:", i)
	fmt.Println("Pointer:", *p1)

	// type, initial size, capacity
	numbers := make([]int, 5, 6)
	numbers[0] = 134
	numbers[1] = 25
	numbers[2] = 33
	numbers[3] = 5234
	numbers[4] = 12
	fmt.Println(numbers)
}
