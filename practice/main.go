package main

import (
	"fmt"
)

/*
 * Reference: Chapter 5 > Lesson 2
 *
 * Define functions as methods of custom types
 */

func main() {
	dog := Dog{
		Name:   "Fido",
		Age:    10,
		Breed:  "Poodle",
		Weight: 45,
		Sound:  "Woof",
	}

	fmt.Println(dog)
	fmt.Printf("%+v\n", dog)

	fmt.Println(dog.Name)
	fmt.Println(dog.Age)
	fmt.Println(dog.Breed)
	fmt.Println(dog.Weight)

	dog.Bark()
	dog.Sound = "Bark"
	dog.Bark()

	dog.Sound = "Arf!"
	dog.BarkTrreeTimes()
}

type Dog struct {
	Name   string
	Age    int
	Breed  string
	Weight int
	Sound  string
}

// Bark is a method of the Dog type
func (d Dog) Bark() {
	fmt.Println(d.Sound)
}

// BarkTrreeTimes is a method of the Dog type
func (d Dog) BarkTrreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}
