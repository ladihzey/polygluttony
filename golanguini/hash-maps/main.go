package main

import (
	"fmt"
	"sort"
)

func main() {
	colors := map[string]string{
		"Red":   "#ff0000",
		"Green": "#4bf745",
		"Blue":  "#0000ff",
	}
	fmt.Println(colors)

	colors["White"] = "#ffffff"
	fmt.Println(colors)

	delete(colors, "Red")
	fmt.Println(colors)
	fmt.Println()

	for k, v := range colors {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}
	fmt.Println()

	keys := make([]string, len(colors))
	i := 0
	for k := range colors {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	for i := range keys {
		fmt.Printf("Key: %v, Value: %v\n", keys[i], colors[keys[i]])
	}
	fmt.Println()
}
