package algs4

import (
	"bufio"
	"log"
	"os"
)

type In struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewIn(path string) In {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	in := In{
		file:    file,
		scanner: bufio.NewScanner(file),
	}
	return in
}

func ScanAll(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return len(data), data, nil
}

func (in In) ReadAll() string {
	in.scanner.Split(ScanAll)
	in.scanner.Scan()
	return in.scanner.Text()
}

func (in In) Close() {
	if in.file != nil {
		in.file.Close()
	}
}
