package main

import "fmt"

func zeroValue(iValue int) {
	iValue = 0
}

func zeroPointer(iPointer *int)  {
	*iPointer = 0
}

func main() {
	i := 1
	fmt.Println("i =", i)

	zeroValue(i)
	fmt.Println("i from function zeroValue =", i)

	zeroPointer(&i)
	fmt.Println("i vaule from function zeroPointer =", i)
	fmt.Println("i address from function zeroPointer =", &i)

	
}