package main

import (
	"fmt"
	"github.com/efuquen/algorithms/pkg/algs4/stdin"
	"github.com/efuquen/algorithms/pkg/ch1"
)

func main() {
	s := ch1.NewStack[string]()

	str, err := stdin.ReadString()
	item := &str
	fmt.Println()
	for err == nil {

		if *item != "-" {
			s.Push(item)
		} else if !s.IsEmpty() {
			popItem := s.Pop()
			fmt.Printf("Pop: %s\n", *popItem)
		}

		str, tmpErr := stdin.ReadString()
		err = tmpErr
		item = &str
	}

	fmt.Printf("(%d left on stack )\n", s.Size())
	i := 0
	for str := range s.Iter() {
		fmt.Printf("Stack[%d]: %s Addr: %p\n", i, *str, str)
		i += 1
	}
}
