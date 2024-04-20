package main

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data1 := []byte("hello\nworld")
	// write file
	err := os.WriteFile("data.txt", data1, 0644)

	check(err)

	// create file
	file, err := os.Create("Employee Name")
	check(err)

	defer file.Close()

	data2 := []byte("Sira\nManee")
	os.WriteFile("EmployeeName.txt", data2, 0644)
}
