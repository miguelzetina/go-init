package main

import (
	"fmt"
	"math"
)

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

	// Rectangle
	var high = 5
	var width = 6
	var result = high * width
	fmt.Println("Rectangle area is:", result)

	// Trapezoid
	var baseA = 38.0
	var baseB = 18.0
	var high2 = 7.0
	var result2 = ((baseA + baseB) / 2.0) * high2
	fmt.Println("Trapezoid area is:", result2)

	// Circle
	var radio = 7.0
	var result3 = math.Pi * math.Pow(radio, 2)
	fmt.Println("Circle area is:", result3)

	// Printf
	name := "Mike"
	age := 26
	fmt.Printf("%s is %d years old\n", name, age)
	fmt.Printf("%v is %v years old\n", name, age)

	// Sprintf
	message := fmt.Sprintf("%v is %v years old", name, age)
	fmt.Println(message)

	// Data type
	fmt.Printf("name type: %T\n", name)
	fmt.Printf("age type: %T\n", age)
}
