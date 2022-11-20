package main

import (
	"fmt"
	"myproject/mylib"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	avg := mylib.Average(s)
	fmt.Println(avg)
}
