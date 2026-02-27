package utils

import (
	"bufio"
	"errors"
	"io"
	"os"
)

type FileReader struct {
	file *os.File
	ChunkSize int

	currLine ReadLine
}

type ReadLine struct {
	scanned bool
	scanner bufio.Scanner

	currIndex int
	line string
}

type FileReaderConfig struct {
	Path string
	ChunkSize int
}

func NewFileReader(config FileReaderConfig) (*FileReader, error) {
	if config.Path == "" {
		return nil, errors.New("Path cannot be empty.")
	}

	file, err := os.Open(config.Path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	readLine := ReadLine{
		scanner: *scanner,
		scanned: false,
	}

	if config.ChunkSize == 0 {
		return &FileReader{file, 1024, readLine}, nil
	}

	return &FileReader{file, config.ChunkSize, readLine}, nil
}

func (fs *FileReader) Next() (char byte, err error) {
	if !fs.currLine.scanned {
		fs.currLine.scanner.Scan()
		fs.currLine.line = fs.currLine.scanner.Text()
		fs.currLine.currIndex = 0
		fs.currLine.scanned = true
		return fs.currLine.line[fs.currLine.currIndex], nil
	}

	if fs.currLine.currIndex == len(fs.currLine.line) - 1 {
		if dataExists := fs.currLine.scanner.Scan(); dataExists {
			fs.currLine.line = fs.currLine.scanner.Text()
			fs.currLine.currIndex = 0
			return fs.currLine.line[len(fs.currLine.line) - 1], nil
		} else {
			return byte(0), io.EOF
		}
	}

	fs.currLine.currIndex++
	return fs.currLine.line[fs.currLine.currIndex], nil
}

func (fs *FileReader) Close() error {
	err := fs.file.Close()

	if err != nil {
		return err
	}
	return nil
}
