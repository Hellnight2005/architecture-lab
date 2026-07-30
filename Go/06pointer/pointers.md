# Pointers in Go

A **pointer** is a variable that stores the **memory address of another variable**.

Pointers are useful when you want to access or modify the original value through its memory address.

---

## Creating a Pointer

A pointer can be declared using `*` followed by the data type:

```go
var ptr *int
```

Here:

- `ptr` is a pointer variable.
- `*int` means `ptr` can store the address of an `int`.
- The default value of an uninitialized pointer is `nil`.

```text
ptr → nil
```

---

## Address-of Operator `&`

The `&` operator is used to get the **memory address** of a variable.

```go
mynumber := 23
ptr := &mynumber
```

Conceptually:

```text
mynumber
   │
   │ value = 23
   ▼
┌─────────┐
│   23    │
└─────────┘
     ▲
     │
     │ address
     │
    ptr
```

`ptr` stores the address of `mynumber`.

---

## Dereference Operator `*`

The `*` operator is used to access the **value stored at the address** held by a pointer.

```go
fmt.Println(*ptr)
```

If `ptr` points to `mynumber`, then:

```go
*ptr
```

gives the value of `mynumber`.

So:

```text
ptr   → memory address
*ptr  → value at that address
```

---

## Modifying a Value Through a Pointer

You can modify the original variable using the pointer:

```go
*ptr = *ptr + 2
```

Since `ptr` points to `mynumber`, changing `*ptr` also changes `mynumber`.

```text
Before:

mynumber = 23
ptr ───────► 23


After:

*ptr = *ptr + 2

mynumber = 25
ptr ───────► 25
```

The pointer does not contain a separate copy of the value. It points to the same memory location.

---

## Important Operators

| Operator | Name                 | Purpose                                            |
| -------- | -------------------- | -------------------------------------------------- |
| `&`      | Address-of operator  | Gets the memory address of a variable              |
| `*`      | Dereference operator | Accesses the value stored at the pointer's address |
| `*T`     | Pointer type         | Declares a pointer to type `T`                     |

Example:

```go
x := 10
ptr := &x
```

```text
x       → value: 10
&x      → address of x
ptr     → stores address of x
*ptr    → value at that address: 10
```

---

## Key Points

- A pointer stores a memory address.
- `&variable` gets the address of a variable.
- `*pointer` accesses the value at that address.
- Modifying `*pointer` modifies the original variable.
- An uninitialized pointer has the zero value `nil`.
- `*int` means a pointer to an `int`.
