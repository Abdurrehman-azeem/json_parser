package json_parser

import (
	"fmt"
	"io"
	"regexp"

	FileReader "github.com/aa/v2/utils"
)

type Token struct {
	TokenType string
	Token string
}

type LexicalAnalyzer struct {
	Tokens []Token
	fileReader *FileReader.FileReader
}

func (la LexicalAnalyzer) Tokenize() error {
	defer la.fileReader.Close()

	_, err := la.fileReader.Next()
	if err != nil {
		return err
	}

	err = la.skipCharacters()
	if err != nil {
		return err
	}

	switch la.fileReader.CurrChar() {
		case byte('{'):
			la.tokenizeObject()
		case byte('['):
			la.tokenizeArray()
	}

	err = la.skipCharacters()
	if err != nil {
		return err
	}

	if la.fileReader.CurrChar() != 0 {
		if err != nil {
			return fmt.Errorf("EOF expected instead got, %c", la.fileReader.CurrChar())
		}
	}

	return nil
}

func (la *LexicalAnalyzer) tokenizeObject() error {
	la.Tokens = append(la.Tokens, Token{
		Token: string(byte('{')),
		TokenType: "LEFT-BRACKET",
	})

	_, err := la.fileReader.Next()
	if err != nil {
		return err
	}

	err = la.skipCharacters()
	if err != nil {
		return err
	}

	if la.fileReader.CurrChar() != byte('"') {
		return fmt.Errorf("Expected a property. Instead got %v.", la.fileReader.CurrChar())
	}

	fmt.Println("Did it run till here?")

	for la.fileReader.CurrChar() == '"' {
		err = la.tokenizeKeyValuePair()
		if err != nil {
			return err
		}

		err = la.skipCharacters()
	 	if err != nil {
			return err
		}
	}

	err = la.skipCharacters()
	if err != nil {
		return err
	}

	char := la.fileReader.CurrChar()
	if char == byte('}') {
		la.Tokens = append(la.Tokens, Token{
			Token: string(byte('}')),
			TokenType: "RIGHT-BRACKET",
		})
	} else {
		return fmt.Errorf("Expected '}' instead got %c", char)
	}
	fmt.Println("checking if this is rnning", string(la.fileReader.CurrChar()), la.Tokens)
	return nil
}

func (la *LexicalAnalyzer) tokenizeArray() error {
	return nil
}

func (la *LexicalAnalyzer) tokenizeKeyValuePair() error {
	err := la.skipCharacters()
	if err != nil {
		return err
	}

	err = la.tokenizeString()
	if err != nil {
		return err
	}

	err = la.skipCharacters()
	if err != nil {
		return err
	}

	err = la.tokenizeColon()
	if err != nil {
		return err
	}

	err = la.skipCharacters()
	if err != nil {
		return err
	}

	err = la.tokenizeValue()
	if err != nil {
		return err
	}

	err = la.skipCharacters()
	if err != nil {
		return err
	}

	err = la.tokenizeComma()
	if err != nil {
		return err
	}

	return nil
}

func (la *LexicalAnalyzer) tokenizeComma() error {
	char := la.fileReader.CurrChar()

	if char != byte(',') {
		return fmt.Errorf("Expected a ',' instead got a %c", char)
	}

	_, err := la.fileReader.Next()
	if err != nil {
		return err
	}

	la.Tokens = append(la.Tokens, Token{
		"COMMA",
		",",
	})

	return nil
}

func (la *LexicalAnalyzer) tokenizeValue() error {
	currChar := la.fileReader.CurrChar()

	var err error

	switch currChar {
		case byte('"'):
			err = la.tokenizeString()
			if err != nil {
				return err
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-':
			err = la.tokenizeNumeric()
			if err != nil {
				return err
			}
	}

	defer func(){
		if err == io.EOF {
			err = fmt.Errorf("Unexpected End-of-File.")
		}
	}()

	return nil
}

func (la *LexicalAnalyzer) tokenizeColon() error {
	token := Token{
		"semi-colon",
		":",
	}

	var err error

	char := la.fileReader.CurrChar()
	if char != byte(':') {
		return fmt.Errorf("Expected colon ':' but got %c instead", char)
	}

	char, err = la.fileReader.Next()
	if err != nil {
		return err
	}

	defer func() {
		if err == io.EOF {
			err = fmt.Errorf("Expected colon ':'.")
		}
	}()

	la.Tokens = append(la.Tokens, token)
	return nil
}

func (la *LexicalAnalyzer) tokenizeString() error {
	strData := make([]byte, 10, 20)
	char, err := la.fileReader.Next()
	if err != nil {
		return err
	}

	if char != byte('"') {
		strData = append(strData, char)
	}

	for char != byte('\\') &&  char != byte('"') {
		char, err = la.fileReader.Next()
		if err != nil {
			return err
		}

		if char != byte('"') {
			strData = append(strData, char)
		}
	}

	_, err = la.fileReader.Next()
	if err != nil {
		return err
	}

	newToken := Token{
		"string",
		string(strData),
	}
	la.Tokens = append(la.Tokens, newToken)

	defer func() {
		if err == io.EOF {
			err = fmt.Errorf("Expected '\"' instead.")
		}
	}()
	return nil
}

func (la *LexicalAnalyzer) skipCharacters() error {
	char := la.fileReader.CurrChar()

	if char != byte(' ') && char != byte('\n') && char != byte('\t') {
		return nil
	}

	var err error

	for char == byte(' ') || char == byte('\n')  || char == byte('\t') {
		char, err = la.fileReader.Next()
		if err != nil {
			return err
		}
	}
	return nil
}

func (la *LexicalAnalyzer) tokenizeNumeric() error {
	pattern := "-?(?:0|[1-9]\\d*)(?:\\.\\d+)?(?:[eE][+-]?\\d+)?"

	numericValue := []byte{la.fileReader.CurrChar()}

	char, err := la.fileReader.Next()
	if err != nil {
		return err
	}

	for (char >= 48 && char <= 57) || char == 45 || char == 69 || char == 101 || char == 46 {
		numericValue = append(numericValue, char)
		char, err = la.fileReader.Next()
		if err != nil {
			return err
		}
	}

	match, err := regexp.Match(pattern, numericValue)
	if err != nil {
		fmt.Printf("Error encountered in regex pattern %v", err)
		return err
	}

	defer func() {
		if err == io.EOF {
			err = fmt.Errorf("Expected a numeric instead got %c.", char)
		}
	}()

	if match {
		la.Tokens = append(la.Tokens, Token{
			TokenType: "NUMERIC",
			Token: string(numericValue),
		})
	} else {
		return fmt.Errorf("The numerical value is incorrect.")
	}

	return nil
}

func NewLexicalAnalyzer(path string, bufferSize int) (*LexicalAnalyzer, error) {
	config := FileReader.FileReaderConfig{
		Path: path,
		ChunkSize: bufferSize,
	}

	fileReader, err := FileReader.NewFileReader(config)
	if err != nil {
		return nil, err
	}

	return &LexicalAnalyzer{fileReader: fileReader}, nil
}
