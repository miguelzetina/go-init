package main

import (
	"fmt"
)

func firstFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func calculateDouble(number int) int {
	return number * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	firstFunction("Hello World!")
	tripleArgument(1, 2, "Mike")
	value := calculateDouble(2)
	fmt.Println(value)
	value1, value2 := doubleReturn(1)
	fmt.Println(value1, value2)
	value3, _ := doubleReturn(1)
	fmt.Println(value3)

}
