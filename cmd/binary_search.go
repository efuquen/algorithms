package main

import (
	"fmt"
	"github.com/efuquen/algorithms/pkg/algs4/in"
	"github.com/efuquen/algorithms/pkg/algs4/stdin"
	"io"
	"log"
	"os"
	"sort"
)

func indexOf(a []int, key int) int {
	lo := 0
	hi := len(a) - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if key < a[mid] {
			hi = mid - 1
		} else if key > a[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	in, err := in.NewIn(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	allowed, err := in.ReadAllInts()
	if err != nil {
		log.Fatal(err)
	}
	sort.Ints(allowed)

	key, err := stdin.ReadInt()
	for err == nil {
		if indexOf(allowed, key) == -1 {
			fmt.Println(key)
		}
		key, err = stdin.ReadInt()
	}
	if err != io.EOF {
		log.Fatal(err)
	}
}
