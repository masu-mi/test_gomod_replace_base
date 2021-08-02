package main

import (
	"fmt"

	"github.com/masu-mi/test_gomod_replace_dep_a"
)

func main() {
	fmt.Println("IN main")
	fmt.Printf("%d\n", test_gomod_replace_dep_a.Add(1, 2))
}
