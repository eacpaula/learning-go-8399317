package main

import (
	"fmt"
	"time"
)

/*
 * Reference: Chapter 2 > Lesson 7
 *
 * We learned how to work with time lib
 */

func main() {
	n := time.Now()
	fmt.Println("I recorded this video at: \n", n)

	t := time.Date(2008, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go launched at: \n", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Mon Nov 10 23:00:00 2008")
	fmt.Printf("The type of parsedTime is %T\n", parsedTime)
}
