package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func scanDirectoryPanic(path string) {
	fmt.Println(path)
	// получаем сегмент, содержимое каталога
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	// перебираем содержимое каталога
	for _, file := range files {
		filePath := filepath.Join(path, file.Name()) // arg1/arg2
		// если втречаем каталог, то вызываем функцию
		if file.IsDir() {
			scanDirectoryPanic(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}

func main() {
	scanDirectoryPanic("hfgo")
}
