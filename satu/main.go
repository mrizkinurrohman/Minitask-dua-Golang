package satu

func SumArray(n []int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}