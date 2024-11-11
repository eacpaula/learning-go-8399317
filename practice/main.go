package main

import (
	"fmt"
)

/*
 * Reference: Chapter 5 > Lesson 1
 *
 * Define and call functions
 */

func main() {
	doSomething()

	result := addValues(1, 2)
	fmt.Println(result)

	total := addAllValues(1, 2, 3, 4, 5)
	fmt.Println(total)

	total, quantity := addAllValues2(1, 2, 3, 4, 5)
	fmt.Println(total, quantity)
}

func doSomething() {
	fmt.Println("Do something")
}

func addValues(a, b int) int {
	return a + b
}

func addAllValues(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func addAllValues2(values ...int) (total int, quantity int) {
	for _, value := range values {
		total += value
	}
	quantity = len(values)
	return
}
