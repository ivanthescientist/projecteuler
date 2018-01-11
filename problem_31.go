package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Project Euler #31: 2 GBP in 1, 2, 5, 10, 20, 50, 100, 200 pence coins")
	startTime := time.Now()
	ways := 0

	for p1 := 0; p1 <= 200; p1++ {
		for p2 := 0; p2 <= 100; p2++ {
			for p5 := 0; p5 <= 40; p5++ {
				for p10 := 0; p10 <= 20; p10++ {
					for p20 := 0; p20 <= 10; p20++ {
						for p50 := 0; p50 <= 4; p50++ {
							for p100 := 0; p100 <= 2; p100++ {
								for p200 := 0; p200 <= 1; p200++ {
									curSum := 0
									curSum += p1 * 1
									curSum += p2 * 2
									curSum += p5 * 5
									curSum += p10 * 10
									curSum += p20 * 20
									curSum += p50 * 50
									curSum += p100 * 100
									curSum += p200 * 200

									if curSum == 200 {
										/**
										Uncomment for debug output, takes a lot of time though.
										fmt.Printf("%dx1 + %dx2 + %dx5 + %dx10 + %dx20 + %dx50 + %dx100 + %dx200\n",
											p1, p2, p5, p10, p20, p50, p100, p200)
										**/
										ways++
									}
								}
							}
						}
					}
				}
			}
		}
	}

	fmt.Printf("There are %d ways to combine 2GBP\n", ways)
	fmt.Printf("Calculations took %.03f seconds", time.Now().Sub(startTime).Seconds())
}
