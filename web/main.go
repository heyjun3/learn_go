package main

import (
	"context"
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	ctx := context.Background()
	runJobs(ctx)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("action=Hello status=run")
	fmt.Fprintf(w, "hello world")
	fmt.Println("action=Hello status=done")
}

func log(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("log test")
		fn(w, r)
	}
}

func patch(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("patch")
		fn(w, r)
	}
}

func StartServer() {
	http.HandleFunc("/", patch(log(Hello)))
	http.ListenAndServe(":8080", nil)
}

func runJobs(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ec := make(chan error)
	done := make(chan struct{})

	for i := 0; i < 10; i++ {
		go func() {
			cmd := exec.CommandContext(ctx, "sleep", "30")
			err := cmd.Run()
			if err != nil {
				ec <- err
			} else {
				done <- struct{}{}
			}
		}()
	}

	for i := 0; i < 10; i++ {
		select {
		case err := <-ec:
			return err
		case <-done:
			return nil
		}
	}
	return nil
}

type Command interface {
    execute(string)
}

type buttonX struct {}

func (b buttonX) execute(s string) {
	fmt.Println(s)
}

type buttonY struct {}

func (b buttonY) execute(s string) {
	fmt.Println(s)
}

type buttonA struct {}

func (b buttonA) execute(s string) {
	fmt.Println(s)
}

type buttonB struct {}

func (b buttonB) execute(s string) {
	fmt.Println(s)
}

func handler(s string) Command {
	if s == "x" {
		return buttonX{}
	} else if s == "y" {
		return buttonY{}
	} else if s == "a" {
		return buttonA{}
	} else if s == "b" {
		return buttonB{}
	} else {
		return nil
	}
}

func test() {
	c := handler("x")
	if c != nil {
		c.execute("test")
	}
}
