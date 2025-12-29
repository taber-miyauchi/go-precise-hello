// This file demonstrates Go interfaces and Sourcegraph Precise Code Navigation.
//
// == Testing Precise Navigation in Sourcegraph ==
//
// 1. Go to Definition (type):
//    Instantly jump to where a type is defined, even across files or packages.
//    "SimpleGreeter" on line 38 → Highlights its definition on line 35
//
// 2. Go to Definition (method):
//    Navigate from a method call to its declaration in the interface.
//    "Greet" on line 54 → Highlights its definition on line 31
//
// 3. Find Implementations (method):
//    See all concrete implementations of an interface method across the codebase.
//    "Greet" on line 31 → Shows implementations on lines 38 and 46
//
// 4. Find Implementations (interface):
//    Discover all types that satisfy an interface — useful for understanding polymorphism.
//    "Greeter" on line 53 → Shows implementations: SimpleGreeter (line 35) and FormalGreeter (line 43)
//
// 5. Find References:
//    Shows call sites rather than interface implementations. Answers "who calls this function?"
//    "PrintGreeting" on line 53 → Shows references on lines 59 and 60
//
// 6. Go to Definition (cross-file):
//    Jump to a function defined in another file within the same package.
//    "VersionedGreeting" on line 61 → Highlights its definition in helpers.go line 6
//
// 7. Find References (cross-file):
//    Find all usages of a symbol defined in another file.
//    "VersionedGreeting" in helpers.go → Shows reference on line 73
//
// 8. Go to Definition (constant):
//    Navigate to constant definitions used within functions.
//    "appVersion" in helpers.go line 7 → Highlights its definition on line 3 of helpers.go

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
	PrintGreeting(SimpleGreeter{}, "world")           // Output: Hello, world
	PrintGreeting(FormalGreeter{}, "world")           // Output: Good day, world
	fmt.Println(VersionedGreeting("versioned world")) // Output: [v1] Hello, versioned world
}
