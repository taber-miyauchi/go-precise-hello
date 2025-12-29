package main

import "fmt"

// Greeter defines an interface for greeting.
type Greeter interface {
	Greet(name string) string
}

// SimpleGreeter implements Greeter with a simple greeting.
type SimpleGreeter struct{}

// Greet returns a greeting for the given name.
func (g SimpleGreeter) Greet(name string) string {
	return "Hello, " + name
}

func main() {
	g := SimpleGreeter{}
	fmt.Println(g.Greet("world"))
}
