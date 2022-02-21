package main

import (
	"fmt"
	"github.com/efuquen/algorithms/pkg/algs4/stdin"
	"github.com/efuquen/algorithms/pkg/ch1"
)

func main() {
	s := ch1.NewResizingArrayStack[string]()
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
	fmt.Printf("(%d) left on stack )\n", s.Size())
	i := 0
	for str := range s.Iter() {
		fmt.Printf("Stack[%d]: %s\n", i, str)
		i += 1
	}
}
