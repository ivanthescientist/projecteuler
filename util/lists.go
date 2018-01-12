package util


func ReverseSublist(list []int, start int, end int) {
	if start < end {
		for i, j := start, end; i < j; i, j = i+1, j-1 {
			list[i], list[j] = list[j], list[i]
		}
	}
}
