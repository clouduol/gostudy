// Employee test a struct feature
package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var emp1 = Employee{ID: 1}
var emp2 = Employee{ID: 2}
var emp3 = Employee{ID: 3}
var emp4 = Employee{ID: 4}

func EmployeeByID1(id int) *Employee {
	return &emp1
}

func EmployeeByID2(id int) Employee {
	return emp2
}

func main() {
	fmt.Printf("%v\n", emp1)
	fmt.Printf("%v\n", emp2)
	fmt.Printf("%v\n", emp3)
	fmt.Printf("%v\n", emp4)
	EmployeeByID1(1).Salary = 100
	//EmployeeByID2(1).Salary = 100
	//tmp := &EmployeeByID2(2)
	tmp := EmployeeByID2(2)
	tmp.Salary = 200
	fmt.Printf("%v\n", emp1)
	fmt.Printf("%v\n", emp2)
	fmt.Printf("%v\n", emp3)
	fmt.Printf("%v\n", emp4)
}
