package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%016b\n", os.O_RDONLY)
	fmt.Printf("%016b\n", os.O_WRONLY)
	fmt.Printf("%016b\n", os.O_RDWR)
	fmt.Printf("%016b\n", os.O_CREATE)
	fmt.Printf("%016b\n", os.O_APPEND)

	println("----------------------------------")
	// т.к. у них всего один бит содержит 1, то это означает,
	// что их удобно объединять побитовым оператором ИЛИ и биты не будут путаться
	fmt.Printf("%016b\n", os.O_WRONLY | os.O_CREATE)
	fmt.Printf("%016b\n", os.O_WRONLY | os.O_CREATE | os.O_APPEND)
}
