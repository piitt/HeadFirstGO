package main

import (
	"fmt"
	"headfirstgo/part_08/magazse"
)

func main() {
	var employee magazse.Employee
	employee.Name = "Joy Carr"
	employee.Salary = 60000
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)
}
