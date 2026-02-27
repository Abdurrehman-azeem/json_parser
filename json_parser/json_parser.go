package json_parser

import (
	"bufio"
	"fmt"
	"os"

	utils "github.com/aa/v2/utils"
)

type ParserOptions struct {
	CurlyBracketCount           int
	SquareBracketCount          int
	EncounteredLeftCurlyBracket bool
	ParsingString               bool
	ParsingNumber               bool
}

type Parser struct {
	Scanner   *bufio.Scanner
	Ancestors []string
	Options   ParserOptions
	Res       map[string]any
	Valid     bool
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
		if err != nil {
			fmt.Println("Error encountered at line ", li, ".\n", err)
			return
		}

		err = ParseLine(line, p)
		if err != nil {
			fmt.Println("Error encountered at line ", li, ".\n", err)
		}
	}
	err := CurlyBracketVerifier(p.Options.CurlyBracketCount)
	if err == nil {
		p.Valid = true
	} else {
		fmt.Println(err)
		p.Valid = false
	}
}

func ParseLine(line string, p *Parser) error {
	filteredText := ""
	for _, char := range line {
		switch char {
		case '{':
			if !p.Options.ParsingString {
				p.Options.CurlyBracketCount++
				continue
			}
			filteredText = filteredText + string(char)
		case '}':
			if !p.Options.ParsingString {
				p.Options.CurlyBracketCount--
				continue
			}
			filteredText = filteredText + string(char)
		case '[':
			p.Options.SquareBracketCount++
		case ']':
			p.Options.SquareBracketCount--
		case '"':
			p.Options.ParsingString = !p.Options.ParsingString
		case '\n':
		case '\t':
			continue
		case ' ':
			if p.Options.ParsingString {
				filteredText = filteredText + string(char)
			}
		case '0':
		case '1':
		case '2':
		case '3':
		case '4':
		case '5':
		case '6':
		case '7':
		case '8':
		case '9':
			fmt.Println("number")
		}
	}
	return nil
}

func CurlyBracketVerifier(count int) error {
	if count > 0 {
		return fmt.Errorf("Error: Missing '}'.")
	} else if count < 0 {
		return fmt.Errorf("Error: Missing '{'.")
	}
	return nil
}
