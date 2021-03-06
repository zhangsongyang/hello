package main

import (
	"fmt"

	"example.com/user/hello/morestrings"
	"example.com/user/hello/util"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fmt.Println(cmp.Equal("he", "he"))
	fmt.Println(util.NewUuid())
}
