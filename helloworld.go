package main

import (
	"fmt"
	"os/user"
	"time"
)

func main() {
	defineQ1()
}

func defineQ1() {
	f := 1.11
	fmt.Println(int(f))
	m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	fmt.Printf("%T %v", m, m)
}

func args(params ...int) {
	fmt.Println(len(params), params)
}


func counter() (func() int) {
	x := 0
	return func() int{
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func add(x int, y int) (int, int){
	return x + y, x * y
}

func create() {
	f := func(x int) {
		fmt.Println("inner func", x)
	}
	fmt.Printf("%T", f)
	f(2)
}

func bytes() {
	b := []byte{72, 73}
	fmt.Println(b)
	fmt.Println(string(b))

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))
}

func dict() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	fmt.Println(m["nothing"])

	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)
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
