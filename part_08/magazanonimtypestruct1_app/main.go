package main

import (
	"fmt"
	"headfirstgo/part_08/magazanonimtypestruct"
)

func main() {
	subscriber := magazanonimtypestruct.Subscriber{Name: "Aman Singh"}
	subscriber.Address.Street = "123 Oak St"
	subscriber.Address.City = "Omaha"
	fmt.Println("Street:", subscriber.Address.Street)
	fmt.Println("City:", subscriber.Address.City)

	employee := magazanonimtypestruct.Employee{Name: "Joy Carr"}
	employee.Address.State = "OR"
	employee.Address.PostalCode = "97222"
	fmt.Println("State:", employee.Address.State)
	fmt.Println("PostalCode:", employee.Address.PostalCode)
}
