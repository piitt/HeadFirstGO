// package hfdtflvote предназначен для чтения данных из файла
package hfdtflvote

import (
	"bufio"
	"os"
)

// GetStrings читает файл построчно и возвращает сегмент строк
func GetStrings(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}
