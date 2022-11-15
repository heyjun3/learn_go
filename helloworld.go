package main

import (
	"fmt"
	"os/user"
	"time"
)

func main() {
	var i int = 100
    var j int = 200
    var p1 *int
    var p2 *int
    p1 = &i
    p2 = &j
    i = *p1 + *p2
    p2 = p1
	fmt.Println(p2, &i)
    j = *p2 + i
    fmt.Println(j)
}

func point() {
	n := 100
	fmt.Println(n)
	fmt.Println(*&n)
	var i *int = &n
	fmt.Println(i)
}

func stmtQ() {
	l := []int{100, 300 ,323, 11, 21, 2, 4, 5, 6}
	var min int
	for i, v := range l{
		if i == 0{
			min = v
			continue
		}
		if min > v{
			min = v
		}

	}
	fmt.Println(min)

	m := map[string]int{
		"apple": 200,
		"banana": 300,
		"grapes": 150,
		"orange": 100,
		"papaya": 1000,
		"kiwi": 90,
	}
	var sum int
	for _, v := range m{
		sum += v
	}
	fmt.Println(sum)
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	sum := 0
	// for ; sum < 10 ; {
	// 	sum += 2
	// 	fmt.Println(sum)
	// }
	for sum < 10 {
		sum += 2
		fmt.Println(sum)
	}
	l := []string{"python", "go", "ruby", "java"}
	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}
	for i, v := range l {
		fmt.Println(i, v)
	}
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m{
		fmt.Println(k, v)
	}
	for k := range m{
		fmt.Println(k)
	}
	for _, v := range m{
		fmt.Println(v)
	}
}

func by2(num int) string {
	if num % 2 == 0{
		return "ok"
	}else{
		return "no"
	}
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
