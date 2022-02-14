package main

import (
	"fmt"
	"github.com/efuquen/algorithms/pkg/algs4"
	"log"
	"os"
)

func main() {
	in, err := algs4.NewIn(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	nums, err := in.ReadAllInts()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nums)
}
