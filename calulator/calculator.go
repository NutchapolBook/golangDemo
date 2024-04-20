package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string, number int64) float64 {
	fmt.Printf("%v %v: ", promt, number)
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
	fmt.Printf("Operator :")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func main() {
	result := float64(0)
	fmt.Printf("Input number of inputs: ")
	input, _ := reader.ReadString('\n')
	inputCount, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		message, _ := fmt.Scanf("%v must number only")
		panic(message)
	}

	var i int64 = 1
	for i <= inputCount {
		if i == 1 {
			value := getInput("Enter Number", i)
			result = value
		} else {
			value := getInput("Enter Number", i)
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
				panic("Wrong operator")
			}
			// fmt.Println("Result is", result)
		}
		i++
	}

	fmt.Printf("Result is %v", result)
}
