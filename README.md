# Go Learning Journey

This repository contains my progress learning the Go programming language, organized by topics.

## Directory Structure

- `basics/` - Fundamental Go concepts
  - Variables and types
  - Control structures (if, for, switch)
  - Basic input/output
  
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

## How to Run Examples

Each directory contains multiple `.go` files demonstrating different concepts. To run an example:

```bash
cd directory_name
go run filename.go
```

## Package Names

Each directory follows these package naming conventions:
- Main programs use `package main`
- Library code uses the directory name as the package name
- Test files use the same package name with `_test` suffix
