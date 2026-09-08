# Go URL Handling — Theory Notes

## 1. The `net/url` Package

Go provides the standard library package `net/url` for parsing, inspecting, escaping, and constructing URLs.

```go
import "net/url"
```

The main type we work with is `url.URL`. A parsed URL is represented as a `*url.URL`.

---

## 2. What Is a URL?

A URL (Uniform Resource Locator) identifies a resource on a network.

Example:

```
https://example.com/users?id=10#profile
```

Breakdown:

```
https://example.com/users?id=10#profile
│      │           │       │      │
│      │           │       │      └── Fragment
│      │           │       └───────── Query
│      │           └───────────────── Path
│      └───────────────────────────── Host
└──────────────────────────────────── Scheme
```

Main components: **Scheme, Host, Path, Query, Fragment**

---

## 3. Parsing a URL

Use `url.Parse()` when you have a URL as a string and want to work with its individual components.

```go
u, err := url.Parse("https://example.com/users?id=10")
if err != nil {
    fmt.Println("Error parsing URL:", err)
    return
}
```

`url.Parse()` returns `(*url.URL, error)`. Always handle the error.

> **Important:** `url.Parse()` does **not** make an HTTP request. It only converts the URL string into a structured `url.URL`.

```
URL String
    ↓
url.Parse()
    ↓
*url.URL
```

---

## 4. Important `url.URL` Fields

```go
u, err := url.Parse("https://example.com/users?id=10#profile")
```

Accessible fields:

```go
u.Scheme
u.Host
u.Path
u.RawQuery
u.Fragment
```

### Scheme

```go
fmt.Println(u.Scheme) // https
```

Common schemes: `http`, `https`, `ftp`

### Host

```go
fmt.Println(u.Host) // example.com
```

May include a port: `example.com:8080`

### Path

```go
fmt.Println(u.Path) // /users/123
```

In backend APIs, `GET /users/123` usually means _"get the user whose ID is 123."_

### RawQuery

```go
fmt.Println(u.RawQuery) // id=123&name=abhijeet
```

Note: this does **not** include the leading `?`.

### Fragment

```go
fmt.Println(u.Fragment) // installation
```

The part after `#`.

---

## 5. Query Parameters

```
https://example.com/search?q=golang&page=2
```

Get all query params:

```go
query := u.Query()
```

Return type: `url.Values`, which is conceptually `map[string][]string`.

### Why a map of slices?

A parameter can appear multiple times:

```
?tag=go&tag=backend&tag=api
```

```go
query["tag"] // ["go", "backend", "api"]
```

### Getting one value

```go
query.Get("q") // "golang"
```

### Getting all values for a key

```go
query["tag"] // [go backend api]
```

### Looping through parameters

```go
for key, values := range query {
    fmt.Println("Key:", key, "Values:", values)
}
```

> Map iteration order is not guaranteed.

---

## 6. Constructing a URL

```go
u := &url.URL{
    Scheme: "https",
    Host:   "example.com",
    Path:   "/users",
}

urlString := u.String()
// https://example.com/users
```

### With a manual query string

```go
u := &url.URL{
    Scheme:   "https",
    Host:     "example.com",
    Path:     "/users",
    RawQuery: "page=2&limit=10",
}
// https://example.com/users?page=2&limit=10
```

### With `url.Values` (preferred, handles encoding)

```go
values := url.Values{}
values.Set("page", "2")
values.Set("limit", "10")

u := &url.URL{
    Scheme:   "https",
    Host:     "example.com",
    Path:     "/users",
    RawQuery: values.Encode(),
}
// https://example.com/users?limit=10&page=2
```

---

## 7. `Set()` vs `Add()`

| Method                    | Behavior                                     |
| ------------------------- | -------------------------------------------- |
| `values.Set("tag", "go")` | **Replaces** any existing value for that key |
| `values.Add("tag", "go")` | **Appends** another value for that key       |

```go
values.Add("tag", "go")
values.Add("tag", "backend")
values.Add("tag", "api")
// tag=go&tag=backend&tag=api
```

**Rule of thumb:** `Set` → replace, `Add` → append.

---

## 8. URL Encoding

Characters like spaces aren't safe in a raw URL. Use `url.Values` + `Encode()` instead of manually building query strings:

```go
values := url.Values{}
values.Set("search", "hello world")

queryString := values.Encode()
// search=hello+world
```

---

## 9. Parsing vs. HTTP Requests — Key Distinction

**Parsing** only analyzes the URL string — no network call:

```
URL String → url.Parse() → Structured URL
```

**HTTP request** actually contacts the server:

```
URL String → http.Get() → Internet → Server → HTTP Response
```

---

## 10. Making an HTTP GET Request

```go
import "net/http"

response, err := http.Get(urlString)
if err != nil {
    fmt.Println("Error:", err)
    return
}
defer response.Body.Close()
```

`defer` schedules `Close()` to run when the surrounding function returns — always close the body to release network resources.

### Status codes

```go
response.StatusCode // e.g. 200, 404, 500
response.Status      // e.g. "200 OK"
```

| Code | Meaning               |
| ---- | --------------------- |
| 200  | OK                    |
| 201  | Created               |
| 400  | Bad Request           |
| 401  | Unauthorized          |
| 403  | Forbidden             |
| 404  | Not Found             |
| 500  | Internal Server Error |

### Reading the response body

```go
import "io"

body, err := io.ReadAll(response.Body)
if err != nil {
    fmt.Println("Error reading response:", err)
    return
}
content := string(body)
fmt.Println(content)
```

### Full request flow

```
URL → Parse URL → Send Request → Receive Response
    → Check Status Code → Read Body → Close Body
```

---

## 11. Path vs. Query Parameters

|                 | Example             | Typical meaning                  |
| --------------- | ------------------- | -------------------------------- |
| Path parameter  | `GET /users/123`    | Get the resource with ID 123     |
| Query parameter | `GET /users?id=123` | Get resources filtered by ID 123 |

Exact semantics depend on the application, but the structural distinction matters when designing REST APIs.

---

## 12. Real Backend Example

```
GET /products?category=shoes&page=2&limit=20
```

```go
u, err := url.Parse(
    "https://example.com/products?category=shoes&page=2&limit=20",
)
if err != nil {
    return
}

query := u.Query()
category := query.Get("category")
page := query.Get("page")
limit := query.Get("limit")

fmt.Println(category, page, limit)
```

---

## 13. Where This Matters in Backend Development

- REST APIs
- HTTP clients
- API integrations
- URL shorteners
- Web scrapers
- Authentication systems / OAuth callbacks
- Webhooks
- Pagination & filtering
- Redirect services
- Microservices & HTTP middleware
- Reverse proxies

---

## 14. Cheat Sheet

| Task                        | Method                        |
| --------------------------- | ----------------------------- |
| Parse a URL                 | `url.Parse()`                 |
| Convert URL back to string  | `u.String()`                  |
| Get query parameters        | `u.Query()`                   |
| Get one query value         | `u.Query().Get("key")`        |
| Set a query value (replace) | `values.Set("key", "value")`  |
| Add another value (append)  | `values.Add("key", "value")`  |
| Encode query parameters     | `values.Encode()`             |
| Make HTTP GET request       | `http.Get()`                  |
| Read response body          | `io.ReadAll()`                |
| Close response body         | `defer response.Body.Close()` |

---

## 15. Quick Mental Model

```
                         URL
                          │
                          ▼
                     url.Parse()
                          │
                          ▼
                       *url.URL
                          │
          ┌───────────────┼───────────────┐
          ▼               ▼               ▼
       Scheme           Host            Path
                                          │
                                          ▼
                                      RawQuery
                                          │
                                          ▼
                                       Query()
                                          │
                                          ▼
                                     url.Values
```

Then for an HTTP request:

```
URL → url.Parse() → *url.URL → http.Get() → http.Response
                                                  ├── StatusCode
                                                  └── Body → io.ReadAll()
```

---

## 16. Key Takeaways

- `url.Parse()` parses a URL string into a `*url.URL` — **no network call**.
- `Scheme`, `Host`, `Path`, `RawQuery`, `Fragment` are the core fields.
- `Query()` returns `url.Values`, effectively `map[string][]string`.
- `Get()` retrieves one value; `Set()` replaces; `Add()` appends.
- `Encode()` safely builds a query string.
- `http.Get()` actually sends a request over the network.
- Always check errors and `defer response.Body.Close()`.
- Check `StatusCode` before processing a response body.

---

## 17. Practice Example

```go
package main

import (
    "fmt"
    "net/url"
)

func main() {
    rawURL := "https://example.com/products?category=shoes&page=2"

    u, err := url.Parse(rawURL)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Scheme:", u.Scheme)
    fmt.Println("Host:", u.Host)
    fmt.Println("Path:", u.Path)

    query := u.Query()
    fmt.Println("Category:", query.Get("category"))
    fmt.Println("Page:", query.Get("page"))
}
```

Expected output:

```
Scheme: https
Host: example.com
Path: /products
Category: shoes
Page: 2
```
