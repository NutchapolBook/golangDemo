package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("indexInfo.csv")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	count := 1
	for scanner.Scan() {
		line := scanner.Text()
		item := strings.Split(line, ",")

		// print all data in line
		// fmt.Printf("Line %d : %s\n", count, line)

		// print Currency
		fmt.Println(item[3])

		count++
	}
}
