package main

import (
	"fmt"
	inlib "github.com/efuquen/algorithms/pkg/algs4/in"
	"github.com/efuquen/algorithms/pkg/ch1"
	"log"
	"os"
)

func main() {
	in, err := inlib.NewIn(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	a, err := in.ReadAllInts()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ch1.ThreeSumCount(a))
}
