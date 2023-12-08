package main

import "fmt"

func main() {
	var color string
	fmt.Println("enter you color:")
	fmt.Scanf("%s", &color)
	switch color {
	case "blue":
		fmt.Println("#000FF")
	case "green":
		fmt.Println("#008000")
	case "pink":
		fmt.Println("#FFC0CB")
	case "yellow":
		fmt.Println("#FFFF00")
	default:
		fmt.Println("invalid color")
	}
}
