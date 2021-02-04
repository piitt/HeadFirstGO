package main

import (
	"fmt"
	"headfirstgo/part_08/magazinnerstruct"
)

func main() {
	subscriber := magazinnerstruct.Subscriber{}
	fmt.Printf("%#v\n", subscriber.HomeAddress)
	subscriber.HomeAddress.PostalCode = "68111"
	fmt.Println("Postal Code:", subscriber.HomeAddress.PostalCode)
}
