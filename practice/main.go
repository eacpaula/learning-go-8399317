package main

import (
	"fmt"
)

/*
 * Reference: Chapter 4 > Lesson 1
 *
 * Base concepts about conditionals
 */

func main() {
	theAnwser := 42
	var result string

	if theAnwser < 0 {
		result = "Less than zero"
	} else if theAnwser > 0 {
		result = "Greater than zero"
	} else {
		result = "Equal to zero"
	}

	fmt.Println(result)

	if x := -42; x < 0 {
		result = "Less than zero"
	} else if x > 0 {
		result = "Greater than zero"
	} else {
		result = "Equal to zero"
	}

	fmt.Println(result)
}
