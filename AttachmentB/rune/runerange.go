package main

import "fmt"

func main() {
	asciiString := "ABCDE" // ASCII занимают 1 сисвол 1байт
	utf8String := "БГДЖИ"  // эти символы юникода занимают 2 байта
	asciiBytes := []byte(asciiString)
	utf8Bytes := []byte(utf8String)

	// обработка каждого байта в сегменте (так делать не надо)
	for index, currentByte := range asciiBytes {
		// преобразование байт в строку
		fmt.Printf("%d: %s\n", index, string(currentByte))
	}
	println("----------------------------------------------")
	// обработка каждого байта в сегменте
	for index, currentByte := range utf8Bytes {
		// преобразование байт в строку
		fmt.Printf("%d: %s\n", index, string(currentByte))
	}

	println("---------------------------------------------")

	// обработка по рунам
	for position, currentRune := range asciiString {
		// position индекс текущего байта
		// преобразование руны в строку
		fmt.Printf("%d: %s\n", position, string(currentRune))
	}
	println("----------------------------------------------")
	for position, currentRune := range utf8String {
		// position индекс текущего байта
		// преобразование руны в строку
		fmt.Printf("%d: %s\n", position, string(currentRune))
	}
}
