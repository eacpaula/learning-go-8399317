package main

import (
	"fmt"
)

/*
 * Reference: Chapter 3 > Lesson 3
 *
 * Store ordered values in array
 */

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	fmt.Println("The colors are:", colors)
	fmt.Println("The first color is:", colors[0])

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("The numbers are:", numbers)

	fmt.Println("Number of colors:", len(colors))
	fmt.Println("Number of numbers:", len(numbers))
}
