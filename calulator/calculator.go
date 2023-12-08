package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string) float64 {
	fmt.Printf("%v", promt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message, _ := fmt.Scanf("%v must number only", promt)
		panic(message)
	}

	return value
}

func add(value1, value2 float64) float64 {
	return value1 + value2
}

func subtract(value1, value2 float64) float64 {
	return value1 - value2
}

func multiply(value1, value2 float64) float64 {
	return value1 * value2
}

func divide(value1, value2 float64) float64 {
	return value1 / value2
}

func getOperator() string {
	fmt.Printf("operator is ( + - * /)")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func main() {
	result := float64(0)
	fmt.Println("Input number of inputs:")
	input, _ := reader.ReadString('\n')
	inputCount, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		message, _ := fmt.Scanf("%v must number only")
		panic(message)
	}

	for i := 0; i < inputCount; i++ {
		value := getInput("value = ")

		switch operator := getOperator(); operator {
		case "+":
			result = add(result, value)
		case "-":
			result = subtract(result, value)
		case "*":
			result = multiply(result, value)
		case "/":
			result = divide(result, value)
		default:
			panic("wrong operator")
		}
	}

	fmt.Printf("result is %v", result)
}
