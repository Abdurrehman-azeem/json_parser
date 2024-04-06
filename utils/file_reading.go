package utils

import (
	"bufio"
)

func ReadLine(scanner *bufio.Scanner) (string, bool, error) {
	linesRemain := scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", false, err
	}
	return scanner.Text(), linesRemain, nil
}

func delimitters(character rune) bool {
	return character == '\t' || character == '\n' || character == ' '
}
