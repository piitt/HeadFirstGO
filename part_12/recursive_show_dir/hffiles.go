package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func scanDirectory(path string) error {
	fmt.Println(path)
	// получаем сегмент, содержимое каталога
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Returning error from scanDirectory(\"%s\") call\n", path)
		return err
	}
	// перебираем содержимое каталога
	for _, file := range files {
		filePath := filepath.Join(path, file.Name()) // arg1/arg2
		// если втречаем каталог, то вызываем функцию
		if file.IsDir() {
			err := scanDirectory(filePath)
			if err != nil {
				fmt.Printf("Returning error from scanDirectory(\"%s\") call\n", path)
				return err
			}
		} else {
			fmt.Println(filePath)
		}
	}
	return nil
}

func main() {
	err := scanDirectory("hfgo")
	if err != nil {
		log.Fatal(err)
	}
}
