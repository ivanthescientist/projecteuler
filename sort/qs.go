package sort

import "fmt"

func main() {
	fmt.Println("Quicksort demo")
	array := []int{3, 2, 5, 1, 3, 4}
	QuickSort(array, 0, 5)

	fmt.Println(array)
}

func QuickSort(array []int, low int, high int) {
	if low >= high {
		return
	}

	pivot := Partition(array, low, high)
	QuickSort(array, low, pivot-1)
	QuickSort(array, pivot+1, high)
}

func Partition(array []int, low int, high int) int {
	pivot := array[high]
	i := low - 1

	for j := low; j < high; j++ {
		if array[j] < pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}

	if array[high] < array[i + 1] {
		array[i + 1], array[high] = array[high], array[i + 1]
	}

	return i + 1
}
