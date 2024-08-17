package main

import "fmt"

func main() {
	poodle := Dog{"Poodle", 34, "Woof"}

	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Println()

	fmt.Printf("Breed: %v, Weight: %v\n", poodle.Breed, poodle.Weight)
	poodle.Weight = 10
	poodle.Breed = "Yorkie"
	fmt.Printf("Breed: %v, Weight: %v\n", poodle.Breed, poodle.Weight)

	poodle.Speak()

	// d passed as a copy, so consequent calls will not modify the caller struct
	poodle.SpeakThreeTimes()
	poodle.SpeakThreeTimes()
}

type Dog struct {
	Breed  string
	Weight int
	Sound string
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func (d Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}
