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
	Scanner                     *bufio.Scanner
	Ancestors                   []string
	CurlyBracketCount           int
	SquareBracketCount          int
	EncounteredLeftCurlyBracket bool
	PreceededByString           bool
	Res                         map[string]any
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
	li := 0
	for line, EOF, err := utils.ReadLine(p.Scanner); EOF; line, EOF, err = utils.ReadLine(p.Scanner) {
		li++
		if err != nil && err != io.EOF {
			utils.EscalateError(fmt.Errorf("Error encountered while reading from file. \n", err))
		}
		for idx, char := range line {
			switch char {
			case '{':
				if p.EncounteredLeftCurlyBracket == true && !p.PreceededByString {
					utils.EscalateError(fmt.Errorf("Invalid '", char, "' encountered at line ", li, " and column ", idx, "."))
				}
				p.CurlyBracketCount++
			case '}':
				p.CurlyBracketCount--
			}
		}

		if p.CurlyBracketCount != 0 {
			utils.EscalateError(fmt.Errorf("Invalid JSON, due to '{' or '}' mismatch."))
		}
	}
}
