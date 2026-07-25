# Memory Management in Go

Go provides **automatic memory management**, so memory allocation and deallocation are handled by the Go runtime and Garbage Collector.

## Memory Allocation

Go provides two commonly used built-in functions for memory allocation and initialization:

### `new()`

`new(T)` allocates memory for a value of type `T` and returns a **pointer** to it.

- Allocates memory for the value.
- The value is initialized to its **zero value**.
- Returns a pointer (`*T`).
- Can be used with any type.

```go
x := new(int)

fmt.Println(*x) // 0
```

> `new()` does not mean the value is uninitialized. The allocated memory is **zero-initialized**.

---

### `make()`

`make()` is used to initialize **slices, maps, and channels**.

- Allocates and initializes the internal data structure.
- Returns the initialized value, **not a pointer**.
- Used only with slices, maps, and channels.

```go
numbers := make([]int, 5)
users := make(map[string]int)
ch := make(chan int)
```

For example:

```go
numbers := make([]int, 5)
```

The slice is initialized with a length of `5`, and its elements have their **zero values**:

```text
[0 0 0 0 0]
```

---

## `new()` vs `make()`

| `new()`                    | `make()`                                   |
| -------------------------- | ------------------------------------------ |
| Allocates a value          | Initializes a data structure               |
| Returns a pointer `*T`     | Returns the initialized value              |
| Zero-initializes the value | Initializes slices, maps, and channels     |
| Works with any type        | Only works with slices, maps, and channels |

```go
p := new(int)        // *int
numbers := make([]int, 5) // []int
```

---

## Garbage Collection (GC)

Go has an **automatic Garbage Collector**.

When allocated memory is no longer reachable by the program, the Garbage Collector identifies it and reclaims the memory automatically.

```text
Memory Allocation
       ↓
Program uses memory
       ↓
Object becomes unreachable
       ↓
Garbage Collector (GC)
       ↓
Memory is reclaimed
```

The programmer normally does not need to manually deallocate memory.

---

## Key Points

- Go uses **automatic memory management**.
- Memory allocation and deallocation are handled automatically.
- `new()` → allocates a zero-valued value and returns a pointer.
- `make()` → initializes slices, maps, and channels.
- `new()` can be used with any type.
- `make()` is only used with slices, maps, and channels.
- Go uses an automatic **Garbage Collector (GC)**.
- The Garbage Collector reclaims memory that is no longer reachable.
- The Go compiler uses **escape analysis** to decide whether values are allocated on the stack or heap.
