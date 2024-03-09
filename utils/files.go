package utils

import (
	"fmt"
	"os"
)

func CreateFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating %s file: %s\n", filename, err)
		return
	}
	defer file.Close()
	fmt.Printf("Created %s file\n", filename)
}

func WriteToFile(filename, data string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening %s file: %s\n", filename, err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(data + "\n"); err != nil {
		fmt.Printf("Error writing to %s file: %s\n", filename, err)
		return
	}
}
