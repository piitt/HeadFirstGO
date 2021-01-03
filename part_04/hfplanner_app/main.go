package main

import (
	"fmt"
	"headfirstgo/part_04/hfdates"
)

func main() {
	days := 3
	fmt.Println("Your appointment is in", days, "days")
	fmt.Println("with a follow-up in", days+hfdates.DaysInWeek, "days")
}
