package mock

type Generator struct{}

func (Generator) Generate(_ interface{}) (string, error) {
	return `package main

import "fmt"

func main() {
    fmt.Println("Hello, World")
}`, nil
}
