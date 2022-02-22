package main

import (
	"fmt"
	"github.com/efuquen/algorithms/pkg/algs4/stdin"
	"github.com/efuquen/algorithms/pkg/ch1"
)

func main() {
	q := ch1.NewQueue[string]()

	str, err := stdin.ReadString()
	item := &str
	fmt.Println()
	for err == nil {

		if *item != "-" {
			q.Enqueue(item)
		} else if !q.IsEmpty() {
			dequeueItem := q.Dequeue()
			fmt.Printf("Dequeue: %s\n", *dequeueItem)
		}

		str, tmpErr := stdin.ReadString()
		err = tmpErr
		item = &str
	}

	fmt.Printf("(%d left on queue)\n", q.Size())
	i := 0
	for str := range q.Iter() {
		fmt.Printf("Queue[%d]: %s Addr: %p\n", i, *str, str)
		i += 1
	}
}
