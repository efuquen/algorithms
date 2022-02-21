package main

import (
	"fmt"
	"github.com/efuquen/algorithms/pkg/ch1"
)

func main() {
	var count int
	fmt.Scan(&count)
	var qf = ch1.NewQuickFind(count)

	var p, q int
	_, err := fmt.Scan(p, q)
	for err == nil {
		if qf.Connected(p, q) {
			continue
		}
		qf.Union(p, q)
		fmt.Printf("%d %d\n", p, q)
		_, err = fmt.Scan(p, q)
	}
}
