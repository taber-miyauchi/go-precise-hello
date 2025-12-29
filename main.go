// This file demonstrates Go interfaces and Sourcegraph Precise Code Navigation.
//
// == Testing Precise Navigation in Sourcegraph ==
//
// 1. Go to Definition (type):
//    Instantly jump to where a type is defined, even across files or packages.
//    Click "SimpleGreeter" on line 59 → jumps to its definition on line 35
//
// 2. Go to Definition (method):
//    Navigate from a method call to its declaration in the interface.
//    Click "Greet" on line 54 → jumps to the interface method on line 31
//
// 3. Find Implementations (method):
//    See all concrete implementations of an interface method across the codebase.
//    Right-click "Greet" on line 31 → shows implementations on lines 38 and 46
//
// 4. Find Implementations (interface):
//    Discover all types that satisfy an interface — useful for understanding polymorphism.
//    Right-click "Greeter" on line 53 → shows SimpleGreeter (line 35) and FormalGreeter (line 43)
//
// 5. Find References:
//    Shows call sites rather than interface implementations. Answers "who calls this function?"
//    Right-click "PrintGreeting" on line 53 → shows usages on lines 59 and 60

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
	PrintGreeting(SimpleGreeter{}, "world") // Output: Hello, world
	PrintGreeting(FormalGreeter{}, "world") // Output: Good day, world
}
