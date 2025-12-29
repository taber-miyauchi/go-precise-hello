# Go Precise Hello

A minimal Go project demonstrating Sourcegraph Precise Code Navigation features.

## Purpose

This lab provides a hands-on example for testing precise navigation capabilities in Sourcegraph, including:

- **Go to Definition** – Jump to type and method definitions
- **Find Implementations** – Discover all concrete implementations of an interface or method
- **Find References** – Locate all call sites of a function

## Structure

- `main.go` – Contains a `Greeter` interface with two implementations (`SimpleGreeter`, `FormalGreeter`) and documented navigation exercises
- `index.scip` – SCIP index file for precise code intelligence

## Generating the SCIP Index

```bash
# Install scip-go
go install github.com/sourcegraph/scip-go/cmd/scip-go@latest

# Generate the index
scip-go

# Clean up (optional)
rm index.scip
```

## Usage

Open this project in Sourcegraph and follow the navigation exercises documented in `main.go` to explore precise code navigation.
