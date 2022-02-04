package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "permission.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Создание файла:", err)
		return
	}
	defer file.Close()

	text := "new line"
	file.WriteString(text)

	if err := os.Chmod(fileName, 0444); err != nil {
		fmt.Println("Права на файл:", err)
		return
	}

	if file, err = os.Open(fileName); err != nil {
		fmt.Println("Открытие файла:", err)
		return
	}
	defer file.Close()

	buf := make([]byte, len(text))
	file.Read(buf)

	fmt.Println(string(buf))

}
