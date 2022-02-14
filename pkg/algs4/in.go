package algs4

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type In struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewIn(path string) (In, error) {
	in := In{}
	file, err := os.Open(path)
	if err != nil {
		return in, err
	}
	in.file = file
	in.scanner = bufio.NewScanner(file)
	return in, err
}

func scanAll(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return len(data), data, nil
}

func (in In) ReadAll() string {
	in.scanner.Split(scanAll)
	in.scanner.Scan()
	return in.scanner.Text()
}

func (in In) ReadAllStrings() []string {
	return strings.Fields(in.ReadAll())
}

func (in In) ReadAllInts() ([]int64, error) {
	ints := []int64{}
	for _, s := range in.ReadAllStrings() {
		num, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		ints = append(ints, num)
	}
	return ints, nil
}

func (in In) Close() {
	if in.file != nil {
		in.file.Close()
	}
}
