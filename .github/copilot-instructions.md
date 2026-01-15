# AI Coding Guidelines for this Go Project

## Project Overview
This is a simple Go CLI application that collects user information (name, age, email) and performs basic number checking logic. It includes an unused math library package for basic arithmetic operations.

## Architecture
- **Main Package**: Entry point in [main.go](main.go) with user input collection and number checking
- **Math Library**: Separate package in [mathlib/](mathlib/) providing basic math functions (Add, Subtract, Multiply, Divide, Modulus)
- **File Organization**: Functions can be split across multiple files within the same package (e.g., checkNumber in [checknumber.go](checknumber.go))

## Key Patterns
- **Number Checking**: Use conditional logic for specific values and parity checks, as seen in `checkNumber()` function
- **User Input**: Collect input using fmt.Scanln with pointers to variables
- **Error Handling**: Basic division by zero check in mathlib.Divide() returns 0 instead of panicking
- **Package Structure**: Keep utility functions in separate files but same package; use subdirectories for distinct packages

## Development Workflow
- **Run Application**: `go run main.go checknumber.go` (since checkNumber is in a separate file)
- **Build**: `go build .` to create executable
- **Module Management**: Use `go mod tidy` for dependency management

## Code Style
- Use strconv.Itoa() for integer to string conversion in print statements
- Comment YouTube links for reference (as in main.go)
- Function naming: camelCase for exported functions in packages, lowercase for internal

## Integration Points
- No external dependencies currently
- Math library designed for potential reuse but not yet integrated into main flow