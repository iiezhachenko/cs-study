package main

import (
	"fmt"
	"cs-study/mit-cs50/common"
	"log"
)

func main() {
	var (
		shift int
		e error
		plain string
		cipher []byte
	)
	fmt.Print("Enter shift: ")
	shift, e = common.StdinInt()
	if e != nil {
		log.Fatal(e)
	}

	fmt.Print("plaintext: ")
	plain, e = common.StdinString()
	if e != nil {
		log.Fatal(e)
	}

	shift = shift % 26

	for _, letter := range plain {
		cipher = append(cipher, byte(int(letter) + shift))
	}
	fmt.Printf("cipher: %s", string(cipher))
}