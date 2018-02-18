package main

import "fmt"

func main() {
	triangles := map[int]int{}

	for a := 1; a <= 500; a++ {
		for b := 1; b <= 500; b++ {
			for c := 1; c <= 500; c++ {
				if a + b + c < 1000 && isTriangle(a, b, c) {
					triangles[a + b + c]++
				}
			}
		}
	}

	maxSum := 0
	maxCount := 0

	for sum, count := range triangles {
		if count > maxCount {
			maxSum = sum
			maxCount = count
		}
		fmt.Printf("Sum: %d Count: %d\n", sum, count)
	}

	fmt.Println(maxSum)
	fmt.Println(maxCount)
}

func isTriangle(a, b, c int) bool{
	return a + b > c && a + c > b && b + c > a
}