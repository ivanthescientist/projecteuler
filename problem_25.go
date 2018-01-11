package main

import (
	"fmt"
	"github.com/ivanthescientist/projecteuler/bignum"
)

func main() {
	fmt.Println("Project Euler #25: 1000-digit Fibonnacci Number")

	first := []int{1}
	second := []int{1}
	index := 3

	for ;; index++ {
		second, first = bignum.Sum(first, second), second
		if len(second) >= 1000 {
			break
		}
	}

	fmt.Println(index)
}
