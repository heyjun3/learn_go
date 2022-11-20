package mylib

func Average(l []int) int {
	sum := 0
	for _, v := range l {
		sum += v
	}
	return (sum / len(l))
}
