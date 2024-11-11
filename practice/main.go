package main

import (
	"fmt"
)

/*
 * Reference: Chapter 3 > Lesson 6
 *
 * Store ordered values in array
 */

func main() {
	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %d\n", poodle.Breed, poodle.Weight)
	poodle.Weight = 15
	fmt.Printf("Breed: %v\nWeight: %d\n", poodle.Breed, poodle.Weight)
}

// Dog struct
type Dog struct {
	Breed  string
	Weight int
}
