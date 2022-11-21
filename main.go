package main

import (
	"context"
	"fmt"
	"regexp"
	"time"
)

func main() {
	run()
}

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func run() {
	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1 * time.Second)
	defer cancel()
	go longProcess(ctx, ch)
	for {
		select {
		case <- ctx.Done():
			fmt.Println(ctx.Err())
			return
		case  <- ch:
			fmt.Println("success")
			return
		}
	}
}

func match() {
	m, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(m)

	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)
}
func timeer() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
