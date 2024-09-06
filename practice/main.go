package main

import (
	"fmt"
)

/*
 * Reference: Chapter 2 > Lesson 2
 *
 * We learned about some type of variables and declarations
 *
 * Declarations
 * - This special type of declaration ":=", only works inside of a function
 * - The default way to declarate in a lot of programming languages is "var" and on Golang is the same, you can use it outside or inside of a function
 * - "Const" declarations could be de
 *
 * Explicity or Implicity
 * - The difference between these type of declarations is the usage of the "type" after the name of the variable
 *
 * Explicity sample:
 * var aRandomVariable int = 64
 *
 * Implicity sample:
 * var anotherRandomVariable = 64
 *
 * Types
 * - The types on Golang have the same behavior of other programming languages.
 */

const aConst int = 64

func main() {
	var aString string = "This is an example of string ;)"
	fmt.Println(aString)
	fmt.Printf("The variable's type is %T\n\n", aString)

	var anInteger int = 42
	fmt.Println(anInteger)
	fmt.Printf("The variable's type is %T\n\n", anInteger)

	var defaultInt int
	fmt.Println(defaultInt)
	fmt.Printf("The variable's type is %T\n\n", defaultInt)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("The variable's type is %T\n\n", anotherString)

	var anotherInt int = 53
	fmt.Println(anotherInt)
	fmt.Printf("The variable's type is %T\n\n", anotherInt)

	myString := "Other example of string with other operator"
	fmt.Println(myString)
	fmt.Printf("The variable's type is %T\n\n", myString)

	fmt.Println(aConst)
	fmt.Printf("The variable's type is %T\n\n", aConst)
}
