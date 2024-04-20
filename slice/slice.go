package main

import "fmt"

var courseName []string

func main() {
	courseName = []string{"Java", "Python"}
	fmt.Println(courseName)
	courseName = append(courseName, "Go", "C", "C#", "HTML", "JavaScript")
	fmt.Println(courseName)

	courseWeb := courseName[4:7]
	fmt.Println(courseWeb)

	courseWeb = courseName[:4]
	fmt.Println(courseWeb)
}
