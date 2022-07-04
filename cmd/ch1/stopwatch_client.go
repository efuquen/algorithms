package main

import (
	"github.com/efuquen/algorithms/pkg/algs4/random"
	"log"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = random.UniformInt(-1000000, 1000000)
	}
}
