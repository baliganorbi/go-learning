[![CodeQL](https://github.com/baliganorbi/go-learning/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/baliganorbi/go-learning/actions/workflows/github-code-scanning/codeql)
![GitHub Release](https://img.shields.io/github/v/release/baliganorbi/go-learning)
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/{owner}/{repo}/badge)](https://scorecard.dev/viewer/?uri=github.com/{owner}/{repo})
[![SLSA 3](https://slsa.dev/images/gh-badge-level3.svg)](https://slsa.dev)

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

## Verifying Build Provenance

The build workflow of this project generates SLSA (Supply-chain Levels for Software Artifacts) provenance. The `verify-provenance.sh` script allows you to verify the authenticity and integrity of the built artifacts.

### Prerequisites
- Install the SLSA verifier tool
- Built artifacts (`program-linux-amd64` and its provenance file `program-linux-amd64.intoto.jsonl`)

### Usage

```bash
# Make the script executable (first time only)
chmod +x verify-provenance.sh

# Verify the provenance
./verify-provenance.sh <source-tag>
```

The script verifies:
- The presence of required files (artifact and provenance)
- The SLSA verifier installation
- The authenticity of the build artifact against its provenance

## Directory Structure

- `basics/` - Fundamental Go concepts
  - Hello World
  - Data types
  - Control structures (if, for, switch)
  
- `functions/` - Working with functions
  - Function declarations
  - Methods
  - Interfaces
  
- `data-structures/` - Data structures in Go
  - Arrays and Slices
  - Maps
  - Structs

- `concurrency/` - Concurrent programming
  - Goroutines
  - Channels
  - Select statements
  
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