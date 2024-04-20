package main

import "fmt"

type employee struct {
	employeeId   string
	employeeName string
	phone        string
}

func main() {
	employeeList := [3]employee{}

	employeeList[0] = employee{
		employeeId:   "101",
		employeeName: "Pradoo",
		phone:        "0900000000",
	}

	employeeList[1] = employee{
		employeeId:   "102",
		employeeName: "Prayad",
		phone:        "0900000001",
	}

	employeeList[2] = employee{
		employeeId:   "103",
		employeeName: "Pranee",
		phone:        "0900000002",
	}

	employeeList2 := []employee{}
	employee1 := employee{
		employeeId:   "101",
		employeeName: "Pradoo",
		phone:        "0900000000",
	}

	employeeList2 = append(employeeList2, employee1)

	fmt.Println("employeeList = ", employeeList)
	fmt.Println("employeeList2 = ", employeeList2)

}
