package utils

import (
	"bufio"
)

// Read File in chunks
func ReadLine(scanner *bufio.Scanner) (string, bool, error) {
	linesRemain := scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", false, err
	}
	return scanner.Text(), linesRemain, nil
}
