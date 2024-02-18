package jsonparser

import (
	"bufio"
	"fmt"
	"os"
)

type ParserInterface interface {
	ParseLine()
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

func (p *Parser) ParseLine() {
	scanner := p.Scanner
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
