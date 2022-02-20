package main

import (
	"fmt"
	"github.com/efuquen/algorithms/pkg/algs4/stdin"
	"github.com/efuquen/algorithms/pkg/ch1"
)

func main() {
	s := ch1.NewFixedCapacityStackOfStrings(100)
	item, err := stdin.ReadString()
	fmt.Println()
	for err == nil {
		if item != "-" {
			fmt.Println("Push:", item)
			s.Push(item)
		} else if !s.IsEmpty() {
			fmt.Println("Pop: ", s.Pop())
		}
		item, err = stdin.ReadString()
	}
	fmt.Println("(", s.Size(), "left on stack )")
}
