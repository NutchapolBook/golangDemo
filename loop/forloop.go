package main

import (
	"fmt"
)

const count = 15

func main() {
	// i := 0
	// for i < 10 {
	// 	fmt.Println("i =", i)
	// 	i++
	// }

	// for j := 0; j < count; j++ {
	// 	fmt.Println("j =", j)
	// }

	var input string
	for {
		fmt.Scanf("%s", &input)
		if input == "exit" {
			break
		}
		fmt.Println("input =", input)
	}
}