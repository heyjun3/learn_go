package main

import (
	// "trading/app/controllers"
	// "trading/config"
	// "trading/utils"
	"fmt"
)

type numbers []int

type Post struct {
	author string
	date string
}

type test interface {
	string | int
}

type values[T test] []T

func testError() error{
	return fmt.Errorf("test Error")
}

func get(name string) (int, error) {
	err := testError()
	if err != nil {
		return 0, fmt.Errorf("fail to get user age with name %s: %w", name, err)
	}
	return 1, nil
}

func (v values[test]) print() {
	for _, n := range v {
		fmt.Println(n)
	}
}

func (nm numbers) Map(f func(int) int) numbers{
	var numbers numbers
	for _, num := range nm {
		numbers = append(numbers, f(num))
	}
	return numbers
}

func (nm numbers) plus(n int) numbers{
	var numbers numbers
	for _, num := range nm {
		numbers = append(numbers, num+n)
	}
	return numbers
}

func (nm numbers) multiple(n int) numbers {
	var numbers numbers
	for _, num := range nm {
		numbers = append(numbers, num * n)
	}
	return numbers
}

func main() {
	// utils.LoggingSettings(config.Config.LogFile)
	// controllers.StreamIngestionData()
	// controllers.StartWebServer()
	tmp()
}

func tmp() {
	_, err := get("Bob")
	if err != nil {
		fmt.Println(err)
	}
}