package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	asciiString := "ABCDE" // ASCII занимают 1 сисвол 1байт
	utf8String := "БГДЖИ"  // эти символы юникода занимают 2 байта

	// длина в байтах
	fmt.Println(len(asciiString)) // 5
	fmt.Println(len(utf8String))  // 10

	// длина строки в символах
	fmt.Println(utf8.RuneCountInString(asciiString)) // 5
	fmt.Println(utf8.RuneCountInString(utf8String))  // 5

	// преобразование строки в сегмент рун
	asciiRunes := []rune(asciiString)
	utf8Runes := []rune(utf8String)
	asciiRunesPartial := asciiRunes[3:]
	utf8RunesPartial := utf8Runes[3:]
	fmt.Println(string(asciiRunesPartial))           // DE
	fmt.Println(string(utf8RunesPartial))            // ЖИ
}
