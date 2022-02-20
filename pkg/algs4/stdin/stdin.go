package stdin

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(bufio.NewReader(os.Stdin))

func init() {
	scanner.Split(bufio.ScanWords)
}

func ReadInt() (int, error) {
	if !scanner.Scan() {
		if scanner.Err() != nil {
			return 0, scanner.Err()
		} else {
			return 0, io.EOF
		}
	}
	num, err := strconv.ParseInt(scanner.Text(), 10, 64)
	return int(num), err
}

func ReadString() (string, error) {
	if !scanner.Scan() {
		if scanner.Err() != nil {
			return "", scanner.Err()
		} else {
			return "", io.EOF
		}
	}
	return strings.TrimSpace(scanner.Text()), nil
}
