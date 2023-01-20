package main


import (
	"fmt"
	"net/http"
)

func main() {
	StartServer()
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