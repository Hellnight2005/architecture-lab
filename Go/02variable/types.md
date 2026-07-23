# Go — Types, Variables, Constants, and Declarations

## 1. Go Type System

Go is a **statically typed** and **strongly typed** programming language.

This means the type of a variable is known and checked at compile time.

For example:

```go
var username string = "hellnight2005"
```

Here, `username` has the type `string`.

You cannot later assign a value of an incompatible type:

```go
username = 123
```

This results in a compile-time error because `123` is an integer and `username` is a string.

---

## 2. Everything Has a Type

In Go, values and expressions have types.

Some common types are:

```text
string
bool
int
uint
float32
float64
```

Go also has more complex types:

```text
array
slice
map
struct
pointer
function
channel
interface
```

---

# 3. Go is Case-Sensitive

Go is a **case-sensitive** language.

This means:

```text
name
Name
NAME
```

are all different identifiers.

For example:

```go
var name string
var Name string
```

These are two different variables.

---

# 4. Exported and Unexported Identifiers

Capitalization has a special meaning in Go.

If an identifier starts with an **uppercase letter**, it is **exported** from the package.

If it starts with a **lowercase letter**, it is **unexported**.

Example:

```go
const LoginToken = "qwertyuiop"
```

`LoginToken` is exported.

```go
const loginToken = "qwertyuiop"
```

`loginToken` is unexported.

### Rule

```text
Uppercase first letter
        ↓
Exported
        ↓
Can be accessed from another package
```

```text
Lowercase first letter
        ↓
Unexported
        ↓
Accessible only within its package
```

This rule applies to identifiers such as:

- Variables
- Constants
- Functions
- Types
- Methods

---

# 5. Strings

The `string` type is used to represent text.

Example:

```go
var username string = "hellnight2005"
```

A string can also be declared using type inference:

```go
var website = "hellnight2005.com"
```

Go determines that:

```text
website → string
```

---

# 6. Boolean

The `bool` type represents a Boolean value.

It has only two possible values:

```text
true
false
```

Example:

```go
var isLoggedIn bool = true
```

---

# 7. Integer Types

Go provides several integer types.

### Signed integers

Signed integers can represent positive and negative values.

```text
int8
int16
int32
int64
```

### Unsigned integers

Unsigned integers can represent only non-negative values.

```text
uint8
uint16
uint32
uint64
```

For example:

```text
uint8
```

can store values from:

```text
0 to 255
```

because it has 8 bits:

```text
2^8 = 256 possible values
```

The values start at `0`, so the maximum value is `255`.

---

## 7.1 `int`

`int` is the general-purpose integer type commonly used in Go.

Its size depends on the architecture:

```text
32-bit architecture → 32 bits
64-bit architecture → 64 bits
```

For general integer calculations, `int` is usually the appropriate default choice.

---

## 7.2 `uintptr`

`uintptr` is an unsigned integer type large enough to hold the bit pattern of a pointer.

It is primarily used in low-level programming and with the `unsafe` package.

It should not normally be used as a general-purpose integer type.

---

# 8. Floating-Point Types

Go provides two floating-point types:

```text
float32
float64
```

`float64` provides greater precision than `float32`.

Example:

```go
var price float64 = 2.356
```

For most general-purpose floating-point calculations, `float64` is commonly used.

---

# 9. Zero Values

Go automatically provides a **zero value** when a variable is declared without an explicit initial value.

Example:

```go
var defaultValue int
```

The zero value of `int` is:

```text
0
```

Common zero values include:

| Type      | Zero Value |
| --------- | ---------- |
| `int`     | `0`        |
| `uint`    | `0`        |
| `float32` | `0`        |
| `float64` | `0`        |
| `bool`    | `false`    |
| `string`  | `""`       |
| Pointer   | `nil`      |
| Slice     | `nil`      |
| Map       | `nil`      |
| Interface | `nil`      |

One of Go's design principles is that variables receive useful default values instead of containing undefined garbage data.

---

# 10. Explicit Variable Declaration

A variable can be declared by explicitly specifying its type.

```go
var username string = "hellnight2005"
```

The structure is:

```text
var + variable name + type + value
```

Example:

```go
var age int = 20
```

Here:

```text
Variable name → age
Type          → int
Value         → 20
```

---

# 11. Type Inference

Go can automatically determine the type of a variable based on the value assigned to it.

Example:

```go
var website = "hellnight2005.com"
```

Go infers:

```text
website → string
```

This is called **type inference**.

The compiler's type-checking process determines the type based on the expression.

The lexer itself does not perform type inference.

A simplified compilation process is:

```text
Source Code
     ↓
Lexer
     ↓
Tokens
     ↓
Parser
     ↓
AST
     ↓
Type Checking / Type Inference
     ↓
Compilation
```

---

# 12. Short Variable Declaration (`:=`)

Go provides a short syntax for declaring and initializing a variable:

```go
numberOfUsers := 30000
```

This is called a **short variable declaration**.

The compiler infers the type:

```text
numberOfUsers → int
```

It is conceptually similar to:

```go
var numberOfUsers int = 30000
```

The `:=` syntax is commonly used inside functions.

Example:

```go
func main() {
    name := "Abhijeet"
    age := 20
}
```

---

# 13. `var` vs `:=`

There are three common ways to declare variables.

### Explicit Type

```go
var name string = "Abhijeet"
```

### Type Inference with `var`

```go
var name = "Abhijeet"
```

### Short Variable Declaration

```go
name := "Abhijeet"
```

In all three cases, Go knows:

```text
name → string
```

The difference is in how the variable is declared.

### Important

The short declaration operator `:=` is generally used **inside functions**.

At package level, use `var` or `const`.

---

# 14. Constants

Constants are values that cannot be changed after declaration.

Example:

```go
const LoginToken string = "qwertyuiop"
```

The value of `LoginToken` cannot be reassigned.

Constants can also be declared with type inference:

```go
const LoginToken = "qwertyuiop"
```

Constants can be declared at package level or inside functions.

Example:

```go
package main

const LoginToken = "qwertyuiop"

func main() {
    // Use LoginToken here
}
```

---

# 15. Unused Variables

Go does not allow a local variable to be declared and then left unused.

For example:

```go
func main() {
    username := "hellnight2005"
}
```

This will produce a compilation error because `username` is declared but never used.

You need to use it:

```go
func main() {
    username := "hellnight2005"

    fmt.Println(username)
}
```

This behavior encourages clean code and prevents unnecessary variables from remaining in the program.

---

# 16. Checking a Variable's Type

The `%T` formatting verb can be used with `fmt.Printf` to display the type of a value.

Example:

```go
fmt.Printf("Variable type is %T\n", username)
```

If `username` is a string, the output is:

```text
Variable type is string
```

Example:

```go
numberOfUsers := 30000

fmt.Printf("Variable type is %T\n", numberOfUsers)
```

Output:

```text
Variable type is int
```

---

# 17. Big Picture

The concepts learned so far can be summarized as:

```text
Go Type System
│
├── Statically Typed
│   └── Types checked at compile time
│
├── Case-Sensitive
│   └── name ≠ Name
│
├── Basic Types
│   ├── string
│   ├── bool
│   ├── int
│   ├── uint
│   ├── float32
│   └── float64
│
├── Composite Types
│   ├── Arrays
│   ├── Slices
│   ├── Maps
│   ├── Structs
│   └── Pointers
│
├── Other Types
│   ├── Functions
│   ├── Channels
│   └── Interfaces
│
├── Variable Declarations
│   ├── var name string = "Abhijeet"
│   ├── var name = "Abhijeet"
│   └── name := "Abhijeet"
│
├── Constants
│   └── const LoginToken = "qwertyuiop"
│
└── Visibility
    ├── Uppercase → Exported
    └── Lowercase → Unexported
```

# Key Takeaways

1. Go is **statically typed**.
2. Go is **case-sensitive**.
3. Variables and expressions have types.
4. Go supports **type inference**.
5. `:=` is the **short variable declaration** syntax.
6. Go provides **zero values** for variables that are declared without initial values.
7. Go has several integer types, including `int`, `int8`, `int64`, `uint8`, and `uint64`.
8. `uint8` has a range of `0–255`.
9. `float64` provides more precision than `float32`.
10. Capitalization determines whether an identifier is **exported** or **unexported**.
11. Constants are declared using `const`.
12. Local variables cannot be declared and left unused.
13. `%T` can be used with `fmt.Printf` to inspect a value's type.
14. Modern Go projects generally use **Go Modules** for dependency and project management.
