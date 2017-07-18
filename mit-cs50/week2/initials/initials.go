package main

import (
	"fmt"
	"log"
	"strings"

	"cs-study/mit-cs50/common"
)

func main() {
	var initials string
	fmt.Print("Enter full name: ")
	fname, e := common.StdinString()
	if e != nil {
		log.Fatal(e)
	}

	words := strings.Fields(fname)

	for _, word := range words {
		initials += strings.ToUpper(string(word[0]))
	}

	fmt.Printf("Initials: %s", initials)
}
