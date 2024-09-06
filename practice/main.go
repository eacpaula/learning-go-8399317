package main

import (
	"fmt"
	"math"
)

/*
 * Reference: Chapter 2 > Lesson 3
 *
 * We learned the behavior of some math calculations with float variables and how to round the values and show them.
 */

func main() {
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println(intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println(floatSum)

	floatSum2 := math.Round(floatSum)
	fmt.Println("the sum is now", floatSum2)

	floatSum3 := math.Round(floatSum*100) / 100
	fmt.Println("the sum is now", floatSum3)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference: %.2f\n", circumference)
}
