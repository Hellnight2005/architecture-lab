# Defer in Go

The `defer` keyword in Go is used to **delay the execution of a function call until the surrounding function returns**.

When Go encounters a `defer` statement, it does not execute that statement immediately. Instead, it schedules the function call to run later, just before the current function finishes.

## Basic Concept

The execution flow can be understood as:

```text
Function starts
      │
      ▼
Normal statements execute
      │
      ▼
defer statements are registered
      │
      ▼
Function reaches the end
      │
      ▼
Deferred statements execute
```

A deferred function call executes when the surrounding function returns.

---

# Multiple Defer Statements

When multiple `defer` statements are used, Go executes them using the **LIFO principle**:

> **Last In, First Out**

This means that the last `defer` statement encountered is the first one to execute.

For example:

```text
defer Statement A
defer Statement B
defer Statement C
```

The execution order will be:

```text
Statement C
Statement B
Statement A
```

## Execution Order in This Example

Inside the `main()` function, the program first executes all normal statements.

Conceptually:

```text
1. Print: Defer!

2. Register:
   This is the third deferred statement.

3. Register:
   This is the first deferred statement.

4. Print:
   This is the second statement.

5. Call myDferFunc()
```

The deferred statements inside `main()` do not execute immediately. They remain registered until `main()` returns.

---

# Defer Inside `myDferFunc()`

The `myDferFunc()` function contains a loop:

```text
for i := 0; i < 5; i++
```

During each iteration, a deferred `fmt.Println(i)` call is registered.

The values are registered in this order:

```text
0
1
2
3
4
```

However, because deferred calls execute using **LIFO**, they are executed in the reverse order when `myDferFunc()` returns:

```text
4
3
2
1
0
```

The flow is:

```text
myDferFunc() starts
        │
        ▼
Register defer: 0
        │
        ▼
Register defer: 1
        │
        ▼
Register defer: 2
        │
        ▼
Register defer: 3
        │
        ▼
Register defer: 4
        │
        ▼
Function returns
        │
        ▼
Execute deferred calls
        │
        ▼
4 → 3 → 2 → 1 → 0
```

After all deferred statements inside `myDferFunc()` are executed, control returns to `main()`.

---

# Final Execution Flow

The complete program executes in the following order:

```text
Defer!

This is the second statement.

4
3
2
1
0

This is the first deferred statement.

This is the third deferred statement.
```

The two deferred statements in `main()` also follow the **LIFO principle**.

They were registered in this order:

```text
1. This is the third deferred statement.
2. This is the first deferred statement.
```

Therefore, they execute in reverse order:

```text
1. This is the first deferred statement.
2. This is the third deferred statement.
```

---

# Important Points About `defer`

- `defer` delays the execution of a function call until the surrounding function returns.
- Normal statements continue executing after a `defer` statement is encountered.
- Multiple deferred calls execute in **Last In, First Out (LIFO)** order.
- Each function manages and executes its own deferred calls when that function returns.
- A deferred call inside `myDferFunc()` executes when `myDferFunc()` returns.
- A deferred call inside `main()` executes when `main()` returns.

## Common Use Cases

The `defer` keyword is commonly used for cleanup operations, such as:

- Closing files.
- Closing database connections.
- Unlocking mutexes.
- Releasing resources.
- Performing cleanup before a function returns.

Using `defer` helps ensure that cleanup code is executed even when a function has multiple return points.

## Key Takeaway

`defer` allows a function call to be postponed until the current function is about to return.

When multiple deferred calls are present, Go executes them in **reverse order of registration**, following the **LIFO (Last In, First Out)** principle.

```text
First registered  → Executes last
Last registered   → Executes first
```
