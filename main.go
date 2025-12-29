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

// FormalGreeter implements Greeter with a formal greeting.
type FormalGreeter struct{}

// Greet returns a formal greeting for the given name.
func (g FormalGreeter) Greet(name string) string {
	return "Good day, " + name
}

// PrintGreeting accepts ANY Greeter and prints the greeting.
// This is where the interface becomes useful - you can pass
// SimpleGreeter, FormalGreeter, or any future implementation.
func PrintGreeting(g Greeter, name string) {
	fmt.Println(g.Greet(name))
}

func main() {
	// Same function, different implementations - that's the power of interfaces
	PrintGreeting(SimpleGreeter{}, "world")  // Output: Hello, world
	PrintGreeting(FormalGreeter{}, "world")  // Output: Good day, world
}
