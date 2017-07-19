package main

import (
	"os"
	"log"
	"strings"
	"fmt"
	"cs-study/mit-cs50/common"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Please provide keyphrase as argument.")
	}

	var (
		keyword string = strings.ToLower(os.Args[1])
		shift int
		plain string
		cipher []byte
		e error
	)

	if len(keyword) == 0 {
		log.Fatal("Keyword can't be empty.")
	}

	fmt.Print("plaintext: ")
	if plain, e = common.StdinString(); e != nil {
		log.Fatal(e)
	}

	for i, letter := range plain {
		shift = int(keyword[i % len(keyword)] - 'a')
		cipher = append(cipher, byte(int(letter) + shift))
	}

	fmt.Printf("cipher: %s", string(cipher))
}
