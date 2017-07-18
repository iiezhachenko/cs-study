package main

import (
	"fmt"
	"log"
)

const bottles_per_min int = 12

func main() {
	var (
		min int
	)
	fmt.Print("How much time did you spent in shower?\nInput: ")
	if _, e := fmt.Scan(&min); e != nil {
		log.Fatalf("Invalid value \"%v\". Should be an integer.\n", min)
	}
	fmt.Printf("You spent %v minutes in shower and spent %d bottles of drinking water.\n",
				min, min * bottles_per_min)
}
