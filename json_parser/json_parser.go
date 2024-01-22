package jsonparser

import (
	"fmt"
	"os"
)

func ParseFromFile(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error encountered while trying to access JSON file.\n", err)
	}
	fmt.Println(file)
	return true
}
