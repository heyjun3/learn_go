package main

import (
	"fmt"
	"os/user"
	"time"
)

func main() {
	list()
}

func hello() {
	fmt.Println("Hello World", time.Now())
	fmt.Println(user.Current())
}
func bool() {
	t, f := true, false
	fmt.Printf("%T %v\n", t, t)
	fmt.Printf("%T %v\n", f, f)
}

func list() {
	var a [2]int
	a[0] = 100
	a[1] = 222
	fmt.Println(a)

	var b [2]int = [2]int{100, 200}
	fmt.Println(b)

	var c []int = []int{300, 400}
	fmt.Println(c)
	var d = append(c, 500)
	fmt.Println(d)

	e := [][]int{
		[]int{1, 2, 3},
		[]int{1, 4, 9},
		[]int{1, 33, 11},
	}
	fmt.Println(e)

	// f := make([]int, 5)
	f := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		f = append(f, i)
		fmt.Println(f)
	}
	fmt.Println(f)
}
