package main

import "fmt"

func calmDown() {
	recover()
}

func freakOut() {
	defer calmDown()
	panic("oh no")                 // точка возврата управления
	fmt.Println("I won't be run!") // этот код не выполнится
}

func main() {
	freakOut()
	fmt.Println("Exiting normally")
}
