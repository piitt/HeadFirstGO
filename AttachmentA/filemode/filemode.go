package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.FileMode(0700))
	fmt.Println(os.FileMode(0070))
	fmt.Println(os.FileMode(0007))

	// FilMode основан на uint32
	println("================")
	fmt.Println(os.FileMode(17))
	fmt.Println(os.FileMode(249))
	fmt.Println(os.FileMode(1000))

}