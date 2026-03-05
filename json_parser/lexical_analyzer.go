package json_parser

import (
	"fmt"
	"io"

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

	char, err := la.fileReader.Next()
	if err != nil {
		return err
	}

	err = la.skipCharacters()
	if err != nil {
		return err
	}

	switch char {
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

	var err error
	char := la.fileReader.CurrChar()

	switch char {
		case byte('"'):
			err = la.tokenizeKeyValuePair()
		case byte('['):
			err = la.tokenizeArray()
	}

	err = la.skipCharacters()
	if err != nil {
		return err
	}

	char = la.fileReader.CurrChar()
	if char == byte('}') {
		la.Tokens = append(la.Tokens, Token{
			Token: string(byte('}')),
			TokenType: "RIGHT-BRACKET",
		})
	} else {
		return fmt.Errorf("Expected '}' instead got %c", char)
	}

	return nil
}

func (la *LexicalAnalyzer) tokenizeArray() error {
	return nil
}

func (la *LexicalAnalyzer) tokenizeKeyValuePair() error {
	err := la.tokenizeString()
	if err != nil {
		return err
	}

	err = la.tokenizeColon()
	if err != nil {
		return err
	}

	err = la.tokenizeValue()
	if err != nil {
		return err
	}

	return nil
}

func (la *LexicalAnalyzer) tokenizeValue() error {
	currChar := la.fileReader.CurrChar()

	err := la.skipCharacters()
	if err != nil {
		return err
	}

	switch currChar {
		case byte('"'):
			la.tokenizeString()
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

	char, err := la.fileReader.Next()
	if err != nil {
		return err
	}

	err = la.skipCharacters()
	if err != nil {
		return err
	}

	if (byte(char) >= 33 && byte(char) < 58) ||(byte(char) > 58 && byte(char) < 127) {
		return fmt.Errorf("Expected colon ':' but got %c instead", char)
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
	currChar := la.fileReader.CurrChar()
	char, err := la.fileReader.Next()
	if err != nil {
		return err
	}

	if char != byte('"') {
		strData = append(strData, currChar)
	}
	currChar = char

	for char != byte('\\') &&  char != byte('"') {
		char, err = la.fileReader.Next()
		if err != nil {
			return err
		}

		if char != byte('"') {
			strData = append(strData, char)
		}
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
	fmt.Println(la.Tokens)
	return nil
}

func (la *LexicalAnalyzer) skipCharacters() error {
	char, err := la.fileReader.Next()

	for err != io.EOF && (char == ' ' || char == '\n'  || char == '\t'){
		char, err = la.fileReader.Next();
		if err != nil {
			return err
		}
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
