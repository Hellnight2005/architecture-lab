# Structs in Go

A **struct** in Go is a custom data type that allows us to group multiple related values into a single type.

For example, information about a user may include their name, email, account status, and age. Instead of storing each value in separate variables, we can group them together using a `User` struct.

## Example Concept

```text
User
├── Name
├── Email
├── Status
└── Age
```

Each field inside a struct can have its own data type.

| Field    | Data Type | Purpose                                       |
| -------- | --------- | --------------------------------------------- |
| `Name`   | `string`  | Stores the user's name                        |
| `Email`  | `string`  | Stores the user's email address               |
| `Status` | `bool`    | Stores whether the user is active or inactive |
| `Age`    | `int`     | Stores the user's age                         |

A struct allows us to represent related information as a **single value**.

---

# Creating a Struct Instance

After defining a struct, we can create a value of that struct type.

For example:

```text
User
├── Name: Rahul
├── Email: rahul@gmail.com
├── Status: true
└── Age: 22
```

This created value is commonly referred to as an **instance** or **value** of the struct.

The variable `user` contains all the information related to one specific user.

Structs are useful when multiple pieces of data logically belong together.

---

# What Is a Method?

A **method** is a function that is associated with a specific type.

In this example, methods are associated with the `User` struct. This allows the `User` type to have both:

- **Data** — such as `Name`, `Email`, `Status`, and `Age`.
- **Behavior** — such as checking the user's status or updating information.

For example, the `User` type can have methods for:

- Getting the user's status.
- Updating the user's email.
- Updating the user's age.
- Printing user information.

Instead of creating completely separate and unrelated functions, methods allow us to associate specific behavior directly with a type.

## Conceptually

```text
User
│
├── Data
│   ├── Name
│   ├── Email
│   ├── Status
│   └── Age
│
└── Methods
    ├── GetStatus()
    └── NewEmail()
```

This makes the code more organized because the data and the operations related to that data are connected through the same type.

---

# Method Receivers

A method is connected to a type using a **receiver**.

The receiver appears between the `func` keyword and the method name.

Conceptually:

```text
func (receiver Type) MethodName() {
    // Method logic
}
```

In the `User` example, `u` is the receiver and `User` is the type:

```text
func (u User) GetStatus() {
    // Access User data
}
```

The receiver allows the method to access the fields of the `User` struct.

For example:

```text
u.Status
u.Email
u.Name
```

There are two main types of method receivers in Go:

1. **Value Receiver**
2. **Pointer Receiver**

The difference between them is important because it determines whether the method works with a **copy of the value** or can modify the **original value**.

The details of value receivers, pointer receivers, and when to use each one follow below.
