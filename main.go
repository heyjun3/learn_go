package main

import (
	"log"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"time"
	"database/sql"

	"golang.org/x/sync/semaphore"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	access()
}

var DbConnection *sql.DB

func access() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person(
				name STRING,
				age INT)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}
}


func run1() {
	ctx := context.TODO()
	var s *semaphore.Weighted = semaphore.NewWeighted(2)
	go longProcess1(ctx, s)
	go longProcess1(ctx, s)
	go longProcess1(ctx, s)
	time.Sleep(5 * time.Second)
}

func longProcess1(ctx context.Context, s *semaphore.Weighted) {
	if err := s.Acquire(ctx, 1); err != nil {
		fmt.Println(err)
		return
	}
	defer s.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func parse() {
	b := []byte(`{"name": "mike", "age": 20, "nicknames": ["a", "b", "c"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	v, _ := json.Marshal(p)
	fmt.Println(string(v))
}

type Person struct {
	Name string
	Age int
	Nicknames []string
}

func http_get() {
	// resp, _ := http.Get("http://example.com")
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	base, _ := url.Parse("http://example.com")
	reference, _ := url.Parse("/test   ?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)
}

func test_defer() {
	defer fmt.Println("test1")
	defer fmt.Println("test2")
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
