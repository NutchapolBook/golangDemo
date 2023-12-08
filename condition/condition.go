package main

import "fmt"

var score int

func calculateGrade(score int){
	if score >= 80 {
		fmt.Println("You have score",score, ":A")
	} else if score >= 70 {
		fmt.Println("You have score",score, ":B")
	} else if score >= 60 {
		fmt.Println("You have score",score, ":C")
	} else if score >= 50 {
		fmt.Println("You have score",score, ":D")
	} else {
		fmt.Println("You have score",score, ":F")
	}
}

func main() {
	fmt.Println("enter your score:")
	fmt.Scanf("%d", &score)
	calculateGrade(score)
}