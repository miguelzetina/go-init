package main

import "fmt"

func main() {
	// Constants declaration
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Variables declaration
	number := 12
	var number1 int = 14
	var number2 int

	fmt.Println(number, number1, number2)

	// Zero values
	var zvalue1 int
	var zvalue2 float64
	var zvalue3 string
	var zvalue4 bool
	fmt.Println(zvalue1, zvalue2, zvalue3, zvalue4)
}
