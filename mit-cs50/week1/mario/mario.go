package main

import (
	"fmt"
	"strings"
	"cs-study/mit-cs50/common"
)

func main() {
	var (
		h int
		e error
	)
	for {
		fmt.Print("Height: ")
		if h, e = common.StdinInt(); e != nil {
			continue
		}
	}

	for i:=0; i <= h; i++ {
		spaces := strings.Repeat(" ", h - i)
		hashes := strings.Repeat("#", i)
		gap := "  "
		fmt.Printf("%s%s%s%s\n", spaces, hashes, gap, hashes)
	}
}
