# Go Learning Journey

This repository contains my progress learning the Go programming language, organized by topics. The project is structured as a single Go module with multiple packages, each focusing on different aspects of Go programming.

## Project Structure

- `main.go` - The main program that demonstrates all examples
- `build_and_run.sh` - Shell script to build and run the program
- `go.mod` - Go module definition file

## Running the Examples

You can run all examples using the build script:

```bash
# Make the script executable (first time only)
chmod +x build_and_run.sh

# Run the examples
./build_and_run.sh
```

Or manually build and run:

```bash
go build -o program main.go
./program
```

## Directory Structure

- `basics/` - Fundamental Go concepts
  - Hello World
  - Data types
  - Control structures (if, for, switch)
  
- `functions/` - Working with functions
  - Function declarations
  - Methods
  - Interfaces
  
- `concurrency/` - Concurrent programming
  - Goroutines
  - Channels
  - Select statements
  
- `data-structures/` - Data structures in Go
  - Arrays and Slices
  - Maps
  - Structs
  
- `packages/` - Package management
  - Creating packages
  - Importing packages
  - Package visibility
  
- `testing/` - Testing in Go
  - Unit tests
  - Benchmarks
  - Examples
  
- `web/` - Web development
  - HTTP servers
  - REST APIs
  - JSON handling
  
- `databases/` - Database interactions
  - SQL
  - NoSQL
  - ORMs
  
- `common/` - Shared utilities and helpers

## Package Organization

Each topic is organized as a separate package:

- Files in topic directories (like `basics/`, `functions/`, etc.) are part of their respective packages
- Each file demonstrates specific concepts and provides reusable functions
- The main program (`main.go`) imports these packages and demonstrates their usage

For example:
```go
import (
    "baliganorbi/learning/basics"    // Import basics package
    "baliganorbi/learning/functions" // Import functions package
)
```