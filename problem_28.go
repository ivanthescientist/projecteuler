package main

import "fmt"

func main() {
	fmt.Println("Project Euler #28: Sum of diagonal's numbers in spiral formed matrix")

	matrix := createMatrix(1001)

	fillMatrixSpiral(matrix)
	//printMatrix(matrix)

	fmt.Printf("Diagonal sum: %d\n", diagonalsSum(matrix))
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, column := range row {
			fmt.Printf("%3d ", column)
		}
		fmt.Println()
	}
}

func createMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			matrix[i][j] = 0
		}
	}

	return matrix
}

func fillMatrixSpiral(matrix [][]int) {
	n := len(matrix)
	cur := n*n
	max := n - 1

	for circle := 0; circle < n - 2; circle++ {
		for i := max - circle; i >= circle; i-- {
			matrix[circle][i] = cur
			cur--
		}

		for j := circle + 1; j <= max-circle; j++ {
			matrix[j][circle] = cur
			cur--
		}

		for i := circle + 1; i <= max-circle; i++ {
			matrix[max-circle][i] = cur
			cur--
		}

		for j := max - circle - 1; j > circle; j-- {
			matrix[j][max-circle] = cur
			cur--
		}
	}
}

func diagonalsSum(matrix [][]int) int {
	sum := 0;
	n := len(matrix)

	for i := 0; i < n; i++ {
		sum += matrix[i][i]
		sum += matrix[i][n-i-1]
	}

	// account for that 1 that happens to be on both diagonals
	return sum - 1
}
