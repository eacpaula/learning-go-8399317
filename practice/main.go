package main

import (
	"fmt"
)

/*
 * Reference: Chapter 4 > Lesson 3
 *
 * Create loops with for statements
 */

func main() {
	colors := []string{"red", "green", "blue", "yellow", "black", "white"}
	fmt.Println(colors)

	// range loop to iterate over a slice of strings and print each element in the slice
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	// range loop to iterate over a slice of strings and print each element in the slice
	for i := range colors {
		fmt.Printf(colors[i])
	}

	// range loop to iterate over a slice of strings and print the index and value of each element in the slice
	for _, color := range colors {
		fmt.Println(color)
	}

	// for loop to iterate over a range of numbers and print each number in the range
	value := 1
	for value < 10 {
		fmt.Println(value)
		value++
	}

	// for loop to iterate over a range of numbers and print each number in the range
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
		if sum > 200 {
			goto theEnd
		}
	}

	// label to jump to
theEnd:
	fmt.Println("The end")

}
