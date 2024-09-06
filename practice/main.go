package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
 * Reference: Chapter 2 > Lesson 3
 *
 * We learned other two different libraries "bufio" and "os"
 *
 * We created a simple example about how to read informations into a terminal application
 */

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered: ", input)
}
