package main

import (
	"fmt"
	test_gomod_replace_dep_a "github.com/masu-mi/test_gomod_replace_dep_a/v2"
	"github.com/masu-mi/test_gomod_replace_dep_a/v2/interfaces"
)

func main() {
	fmt.Println("IN main")
	fmt.Printf("%d\n", test_gomod_replace_dep_a.Add(1, 2))
	fmt.Printf("%d\n", test_gomod_replace_dep_a.Add((&sample{}).Write("100000"), 2))
}

type sample struct{}
var _ interfaces.Writer = &sample{}

func (s *sample) Write(payload string) (num test_gomod_replace_dep_a.OriginalInt) {
	return test_gomod_replace_dep_a.OriginalInt(len(payload))
}
