package utils

import (
	"bufio"
	"fmt"
	"strings"
)

func ReadLine(scanner *bufio.Scanner) (string, error) {
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	fmt.Println(scanner.Text())
	return scanner.Text(), nil
}

func Tokenize(line string) []string {
	return strings.FieldsFunc(line, delimitters)
}

func delimitters(character rune) bool {
	return character == '\t' || character == '\n' || character == ' '
}
