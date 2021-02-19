package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i <= 19; i++ {
		fmt.Printf("%3d: %04o\n", i, i)
	}
	println("=======================================")
	fmt.Printf("Decimal   1: %3d Octal   01: %2d\n", 1, 01)
	fmt.Printf("Decimal  10: %3d Octal  010: %2d\n", 10, 010)
	fmt.Printf("Decimal 100: %3d Octal 0100: %2d\n", 100, 0100)

	println("============================================")
	// цифра восьмиричного числа может быть представленна 3 битами памяти
	// 3 бита необходимо для хранения разрешений одного класса пользователей
	fmt.Printf("%09b\n", 0007)
	fmt.Printf("%09b\n", 0070)
	fmt.Printf("%09b\n", 0700)
	println("=============================================")
	// если бит в двоичном представлении равен 1, то соответствующее разрешение включено
	fmt.Printf("%09b %s\n", 0000, os.FileMode(0000))
	fmt.Printf("%09b %s\n", 0111, os.FileMode(0111))
	fmt.Printf("%09b %s\n", 0222, os.FileMode(0222))
	fmt.Printf("%09b %s\n", 0333, os.FileMode(0333))
	fmt.Printf("%09b %s\n", 0444, os.FileMode(0444))
	fmt.Printf("%09b %s\n", 0555, os.FileMode(0555))
	fmt.Printf("%09b %s\n", 0666, os.FileMode(0666))
	fmt.Printf("%09b %s\n", 0777, os.FileMode(0777))
}
