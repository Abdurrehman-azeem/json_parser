package json_parser

import (
	FileReader "github.com/aa/v2/utils"
)

type Token struct {
	tokenType string
	token byte
}

type LexicalAnalyzer struct {
	tokens []Token
	fileReader *FileReader.FileReader
}

func NewLexicalAnalyzer(path string, bufferSize int) (*LexicalAnalyzer, error) {
	config := FileReader.FileReaderConfig{
		path,
		bufferSize,
	}
	fileReader, err := FileReader.NewFileReader(config)
	if err != nil {
		return nil, err
	}

	return &LexicalAnalyzer{fileReader: fileReader}, nil
}
