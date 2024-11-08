package main

import (
	"fmt"
	"sort"
)

/*
 * Reference: Chapter 3 > Lesson 3
 *
 * Store ordered values in array
 */

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println("The colors are:", colors)

	colors = append(colors, "Yellow")
	fmt.Println("The colors are:", colors)

	colors = append(colors[1:len(colors)])
	fmt.Println("The colors are:", colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println("The colors are:", colors)

	number := make([]int, 5)
	number[0] = 1
	number[1] = 2
	number[2] = 3
	number[3] = 4
	number[4] = 5
	fmt.Println("The numbers are:", number)

	number = append(number, 6)
	fmt.Println("The numbers are:", number)

	sort.Sort(sort.Reverse(sort.IntSlice(number)))
	fmt.Println("The numbers in descending order are:", number)
}
