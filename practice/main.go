package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
 * Reference: Chapter 4 > Lesson 2
 *
 * Evaluate expressions with switch statements
 */

func main() {
	t := time.Now().Unix()
	fmt.Println("Today is", t)

	r := rand.New(rand.NewSource(t))

	var result string

	switch dow := r.Intn(7) + 1; dow {
	case 1:
		result = "Sunday"
		// fallthrough
	case 2:
		result = "Monday"
		// fallthrough
	case 3:
		result = "Tuesday"
		// fallthrough
	case 4:
		result = "Wednesday"
		// fallthrough
	case 5:
		result = "Thursday"
		// fallthrough
	case 6:
		result = "Friday"
		// fallthrough
	case 7:
		result = "Saturday"
		// fallthrough
	default:
		result = "Unknown"
	}

	fmt.Println("Today is", result)
}
