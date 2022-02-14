package cmd

import (
	"fmt"
	"os"

	"github.com/efuquen/algorithms/pkg/algs4/in"
)

func main() {
	in := in.NewIn(os.Args[0])
	fmt.Println(in.ReadAll())
}
