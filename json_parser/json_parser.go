package jsonparser

import (
	"bufio"
	"fmt"
	"io"
	"os"

	utils "github.com/aa/v2/utils"
)

type ParserInterface interface {
	ReadLine()
}

type Parser struct {
	Scanner   *bufio.Scanner
	Ancestors []string
}

func NewParser(path string) Parser {
	fileReader, err := os.Open(path)
	if err != nil {
		fmt.Errorf("Error encountered trying to access file, ", path, ". \n", err)
		return Parser{}
	}

	parser := Parser{
		Scanner:   bufio.NewScanner(fileReader),
		Ancestors: []string{},
	}

	return parser
}

func (p *Parser) ParseFromReader() {
	for line, err := utils.ReadLine(p.Scanner); err == nil; line, err = utils.ReadLine(p.Scanner) {
		if err != nil && err != io.EOF {
			fmt.Errorf("Error encountered while reading from file. \n", err)
		}
		tokens := utils.Tokenize(line)
		for _, token := range tokens {
			fmt.Println(token)
		}
	}
}
