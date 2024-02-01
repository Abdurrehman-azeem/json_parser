package jsonparser

import (
	"fmt"
	"os"
	"bufio"
)

func ParseFromFile(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error encountered while trying to access JSON file.\n", err)
	}
	fmt.Println(file)
	return true
}

func Validate(file.)
