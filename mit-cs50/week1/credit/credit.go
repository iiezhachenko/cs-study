package main

import (
	"fmt"
	"cs-study/mit-cs50/common"
	"log"
)

func luhnAlgo(num string) bool {
	var total int64 = 0
	total += int64(num[len(num)-1] - '0')
	for i:=2; len(num)-i >= 0; i++ {
		digit := int64(num[len(num) - i] - '0')
		if i % 2 == 0 {
			digit *=  2
			if digit > 9 {
				digit -= 9
			}
			total += digit
		} else {
			total += int64(num[len(num) - i] - '0')
		}
	}
	return total % 10 == 0
}

func main() {
	var (
		num string
		e error
	)
	fmt.Printf("Enter card number: ")
	if num, e = common.StdinString(); e != nil {
		log.Fatal(e)
	}
	if !luhnAlgo(num) {
		fmt.Println("INVALID")
		return
	}
	if int(num[0]-'0') == 4 {
		fmt.Println("VISA")
		return
	}

	switch cardType := string(num[0:2]); cardType {
	case "34", "37":
		fmt.Println("AMEX")
		return
	case "51", "52", "53", "54", "55":
		fmt.Println("MASTERCARD")
		return
	default:
		fmt.Println("INVALID")
	}
}