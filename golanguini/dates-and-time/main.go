package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println("Time:", n)

	t := time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("Date:", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2027")
	fmt.Printf("Parsed Time: %v\n", parsedTime)
}
