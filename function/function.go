package main

import "fmt"

func hello() {
	fmt.Println("Hello world")
}

func plus(value1 int, value2 int) int  {
	return value1 + value2
}

func plusThreeValue(value1, value2, value3 int) int {
	return value1 + value2 + value3
}

func main() {
	hello()
	result := plus(1,2)
	fmt.Println("result =", result)

	result2 := plusThreeValue(1,2,3)
	fmt.Println("result2 =", result2)
} 