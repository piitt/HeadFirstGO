package main

import (
	"fmt"
	"headfirstgo/part_08/magazsea"
)

func main() {
	var address magazsea.Address
	address.Street = "123 Oak St"
	address.City = "Omaha"
	address.State = "NE"
	address.PostalCode = "68111"
	fmt.Println(address)
}
