# Go (Golang) — Core Concepts & First Notes

## 1. What is Go?

**Go**, also called **Golang**, is a statically typed, compiled programming language created at Google.

It was designed with a focus on:

- Simplicity
- Fast compilation
- Efficient execution
- Concurrency
- Strong tooling
- Large-scale software development
- Cloud and distributed systems

Go is widely used for:

- System applications
- Backend services
- Web APIs
- Networking software
- Distributed systems
- Cloud infrastructure
- DevOps tools
- Microservices
- CLI tools

A lot of modern cloud infrastructure is built with Go or uses Go heavily.

Examples include technologies such as Kubernetes, Docker, and many cloud-native tools.

---

# 2. Go is a Compiled Language

Go source code is compiled into **native machine code**.

The general flow is:

```text
Go Source Code
      ↓
Go Compiler
      ↓
Machine Code / Executable
      ↓
Operating System
      ↓
CPU
```

For example:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

You can run it directly during development:

```bash
go run main.go
```

However, `go run` is essentially a convenience command that **compiles and executes** the program.

You can also explicitly build an executable:

```bash
go build
```

This produces a binary executable that can be run without requiring a Go virtual machine.

### Important distinction

Go does **not** require a JVM like Java.

```text
Java:
Java Source
    ↓
Compiler
    ↓
Bytecode
    ↓
JVM
    ↓
Machine Code
```

```text
Go:
Go Source
    ↓
Go Compiler
    ↓
Native Executable
    ↓
Operating System
```

Go programs are generally compiled separately for different operating systems and CPU architectures.

For example:

```text
Linux binary      → Linux
Windows binary    → Windows
macOS binary      → macOS
```

Go also has strong support for cross-compilation.

Example:

```bash
GOOS=linux GOARCH=amd64 go build
```

This can produce a Linux executable from another development environment.

---

# 3. What Can Go Be Used For?

Go is a general-purpose language, but it is particularly strong for software that needs performance, concurrency, networking, and operational simplicity.

### Common areas

```text
Go
├── System Applications
├── CLI Tools
├── Web Servers
├── REST APIs
├── gRPC Services
├── Microservices
├── Distributed Systems
├── Networking
├── Cloud Infrastructure
├── DevOps Tools
└── Containers
```

For your interests in **backend engineering, distributed systems, cloud infrastructure, and AI systems**, Go is particularly useful.

A possible architecture could look like:

```text
Client
   ↓
API Gateway
   ↓
Go Backend Service
   ↓
Kafka / Message Queue
   ↓
Worker Services
   ↓
Database / Object Storage
```

---

# 4. "Don't Bring Baggage"

One of the major design philosophies behind Go is to avoid unnecessary complexity.

The language was designed to be relatively small and straightforward.

The idea is:

> Prefer simple language features that are easy to understand, maintain, compile, and use at scale.

Go intentionally does not try to include every feature found in other programming languages.

This means that developers coming from languages with many abstractions may initially find Go restrictive.

However, this simplicity is intentional.

The philosophy is approximately:

```text
Less Language Complexity
        ↓
Easier to Learn
        ↓
Easier to Read
        ↓
Easier to Maintain
        ↓
Easier for Teams to Collaborate
```

---

# 5. Is Go Object-Oriented?

The best answer is:

> **Go is not traditionally object-oriented, but it supports many object-oriented design concepts.**

Go does **not** have traditional classes and inheritance like Java or C++.

Instead, Go uses:

- `struct`
- Methods
- Interfaces
- Composition

Example:

```go
type User struct {
    Name string
}

func (u User) SayHello() {
    fmt.Println("Hello", u.Name)
}
```

Here, `User` is a `struct`, and `SayHello()` is a method associated with it.

Go prefers **composition over inheritance**.

Instead of:

```text
Class A
   ↓ inherits
Class B
   ↓ inherits
Class C
```

Go encourages:

```text
Component A
     +
Component B
     +
Component C
     ↓
Composed Type
```

This is one of the important ideas to understand when coming from Java or C++.

---

# 6. "What You See Is What the Code Does"

Go strongly emphasizes **readability and explicitness**.

The language tries to minimize "magic" behavior.

When reading Go code, you should generally be able to understand what is happening directly from the code.

For example:

```go
result, err := doSomething()

if err != nil {
    return err
}
```

The error handling is explicit.

There is no hidden exception mechanism controlling the normal flow.

This makes Go code relatively predictable and easy to inspect.

The philosophy is:

```text
Readable Code
      ↓
Predictable Behavior
      ↓
Easy Debugging
      ↓
Easy Maintenance
```

This is especially valuable in backend and infrastructure code.

---

# 7. Error Handling — No Traditional try/catch

Go does not use the traditional:

```text
try
catch
finally
```

exception-handling model for ordinary error handling.

Instead, functions commonly return an error value.

Example:

```go
result, err := doSomething()

if err != nil {
    return err
}
```

The pattern is:

```text
Function Call
      ↓
Return Value + Error
      ↓
Check Error
      ↓
Handle Error
```

Example:

```go
file, err := os.Open("data.txt")

if err != nil {
    fmt.Println("Could not open file:", err)
    return
}

defer file.Close()
```

The important concept is:

> Errors are values that the programmer explicitly handles.

Go does have `panic` and `recover`, but these are **not intended to replace normal error handling**.

Generally:

```text
Expected Operational Error
        ↓
Return error

Unexpected / Unrecoverable Condition
        ↓
panic
```

---

# 8. The Compiler and Tooling Do a Lot of Work

Go provides a strong standard toolchain.

Important commands include:

```bash
go run
go build
go test
go fmt
go vet
go mod
```

The Go ecosystem strongly encourages standardized tooling.

For example:

```bash
go fmt
```

automatically formats Go source code.

This means teams don't have to spend much time debating code formatting.

The philosophy is:

```text
Standard Tooling
      ↓
Consistent Code
      ↓
Less Configuration
      ↓
Less Team Friction
```

---

# 9. A Note About the Lexer

Your original note says:

> "lexer does a lot of work"

A more precise way to understand this is that the **Go compiler's frontend** handles several stages of processing your source code.

Conceptually:

```text
Source Code
    ↓
Lexing
    ↓
Tokens
    ↓
Parsing
    ↓
AST
    ↓
Type Checking
    ↓
Intermediate Representation
    ↓
Optimization
    ↓
Machine Code
```

The **lexer** specifically converts source code into tokens.

For example:

```go
x := 10
```

Conceptually becomes tokens such as:

```text
IDENTIFIER: x
OPERATOR: :=
INTEGER: 10
```

The parser then uses those tokens to understand the structure of the program.

So the more accurate statement is:

> **The Go compiler performs a lot of work automatically, while the language itself keeps the syntax and semantics relatively simple.**

---

# 10. The Core Go Philosophy

The ideas you noted can be summarized as:

```text
Go
│
├── Compiled
│     └── Produces native executables
│
├── Simple
│     └── Small language with fewer abstractions
│
├── Explicit
│     └── Code behavior is generally easy to see
│
├── Concurrent
│     └── Goroutines + Channels
│
├── Strongly Typed
│     └── Many errors caught at compile time
│
├── Garbage Collected
│     └── Automatic memory management
│
├── Composition-Oriented
│     └── Structs + Interfaces + Composition
│
├── Cross-Platform
│     └── Strong cross-compilation support
│
└── Cloud-Friendly
      └── Excellent for services and infrastructure
```

# 11. The Most Important Mental Model

When learning Go, don't think:

> "How do I write Java/C++/Python code using Go syntax?"

Instead, think:

> "How does Go want me to design this program?"

The key concepts to focus on next are:

1. **Variables and constants**
2. **Basic types**
3. **Arrays and slices**
4. **Maps**
5. **Structs**
6. **Pointers**
7. **Functions**
8. **Methods**
9. **Interfaces**
10. **Error handling**
11. **Packages**
12. **Modules (`go.mod`)**
13. **Goroutines**
14. **Channels**
15. **Context**
16. **Testing**
17. **HTTP servers**
18. **Database connectivity**
19. **gRPC**
20. **Concurrency patterns**

For your backend and distributed-systems goals, the most important transition will be:

```text
Go Syntax
    ↓
Go Memory Model
    ↓
Pointers + Structs
    ↓
Interfaces + Composition
    ↓
Error Handling
    ↓
Goroutines
    ↓
Channels
    ↓
Concurrency Patterns
    ↓
HTTP / gRPC
    ↓
Distributed Systems
```

That progression will give you a much stronger foundation than simply learning Go syntax and immediately jumping into frameworks.
