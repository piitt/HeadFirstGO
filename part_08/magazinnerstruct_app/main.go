// создание вложенной структуры
package main

import (
	"fmt"
	"headfirstgo/part_08/magazinnerstruct"
)

func main() {
	address := magazinnerstruct.Address{Street: "123 Oak St", City: "Omaha",
		State: "NE", PostalCode: "68111"}
	subscriber := magazinnerstruct.Subscriber{Name: "Aman Singh"}
	subscriber.HomeAddress = address
	fmt.Println(subscriber.HomeAddress)
}
