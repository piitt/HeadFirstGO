package main

import "fmt"

func main() {
	/* избегайте замещение имен
	var int int = 12
	var append string = "minutes of bonus footage"
	var fmt string = "DVD"

	var count int
	var languages = append([]string{}, "Espanol")
	fmt.Println(int, append, "on", fmt, languages)
	*/
	var count int = 12
	var suffix string = "minutes of bonus footage"
	var format string = "DVD"
	var languages = append([]string{}, "Espanol")
	fmt.Println(count, suffix, "on", format, languages)
}
