package mock

type Generator struct {
	GenerateFunc func(interface{}) (string, error)
}

func (g *Generator) Generate(input interface{}) (string, error) {
	if g.GenerateFunc != nil {
		return g.GenerateFunc(input)
	}
	// Default implementation
	return `package main

import "fmt"

func main() {
    fmt.Println("Hello, World")
}`, nil
}
