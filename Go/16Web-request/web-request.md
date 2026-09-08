# Go HTTP Requests — Theory Notes

## 1. The `net/http` Package

Go provides the standard library package `net/http` for working with HTTP.

```go
import "net/http"
```

It can be used to:

- Send HTTP requests
- Create HTTP servers
- Handle HTTP responses
- Work with HTTP headers
- Handle REST APIs
- Build API clients
- Build web applications

For making a simple GET request, the most basic function is `http.Get()`.

---

## 2. What Is an HTTP Request?

When your Go program communicates with a website or API, it sends an HTTP request.

```
Go Program
    |
    | HTTP GET Request
    ↓
Web Server
    |
    | HTTP Response
    ↓
Go Program
```

A **request** can contain:

- HTTP method
- URL
- Headers
- Request body

A **response** can contain:

- Status code
- Headers
- Response body

---

## 3. HTTP Methods

The most common HTTP methods:

| Method   | Purpose                              | Example             |
| -------- | ------------------------------------ | ------------------- |
| `GET`    | Retrieve data                        | `GET /users`        |
| `POST`   | Create/send data                     | `POST /users`       |
| `PUT`    | Completely replace/update a resource | `PUT /users/123`    |
| `PATCH`  | Partially update a resource          | `PATCH /users/123`  |
| `DELETE` | Delete a resource                    | `DELETE /users/123` |

For this lesson, we're mainly working with **GET**.

---

## 4. Making a GET Request

The simplest way to make a GET request in Go:

```go
response, err := http.Get(url)
```

Example:

```go
response, err := http.Get("https://example.com")
```

The function returns `(*http.Response, error)` — so we need to handle the error.

---

## 5. Handling the Error

Always check the error returned from `http.Get()`:

```go
response, err := http.Get(url)
if err != nil {
    fmt.Println("Error:", err)
    return
}
```

The request can fail due to:

- No internet connection
- DNS failure
- Invalid URL
- Server unavailable
- Connection timeout
- TLS/HTTPS problems
- Network errors

Never assume the request succeeded.

---

## 6. What Is `http.Response`?

After a successful request, `response` contains information about the server's response.

Its type is `*http.Response`. You can verify it:

```go
fmt.Printf("Response type: %T\n", response)
// Response type: *http.Response
```

---

## 7. Important `http.Response` Fields

```go
response.Status
response.StatusCode
response.ContentLength
response.Header
response.Body
```

---

## 8. Response Status

```go
fmt.Println("Response:", response.Status)
// 200 OK
```

`Status` contains both the status code and its description.

---

## 9. HTTP Status Code

Get just the numeric code:

```go
fmt.Println("Status Code:", response.StatusCode)
```

| Code | Meaning               |
| ---- | --------------------- |
| 200  | OK                    |
| 201  | Created               |
| 204  | No Content            |
| 301  | Moved Permanently     |
| 302  | Found                 |
| 400  | Bad Request           |
| 401  | Unauthorized          |
| 403  | Forbidden             |
| 404  | Not Found             |
| 500  | Internal Server Error |
| 502  | Bad Gateway           |
| 503  | Service Unavailable   |

---

## 10. Why Status Codes Matter

A successful HTTP request **does not** necessarily mean the server returned successful application data.

```
GET /users    → 200 OK
GET /unknown  → 404 Not Found
```

The request may complete successfully at the network level, but the HTTP operation itself can still be unsuccessful. Always inspect `response.StatusCode`.

---

## 11. Response Content Length

```go
fmt.Println("Response content length:", response.ContentLength)
```

This represents the expected body length, when known.

> **Important:** A value of `-1` means the length is unknown, or the body is being sent using a streaming/chunked mechanism. Don't assume `ContentLength` always holds a valid positive size.

---

## 12. Response Headers

```go
fmt.Println(response.Header)

contentType := response.Header.Get("Content-Type")
fmt.Println("Content-Type:", contentType)
```

Common response headers:

- `Content-Type`
- `Content-Length`
- `Cache-Control`
- `Server`
- `Date`
- `Set-Cookie`

---

## 13. Response Body

The actual data returned by the server is in `response.Body`, of type `io.ReadCloser`.

> The body is **not** simply a string — it's a stream of data that you read.

---

## 14. Why Is Body a Stream?

The response body could be HTML, JSON, XML, plain text, image data, video data, or a large file. Because it could be very large, Go represents it as a stream.

`response.Body` implements both:

- `io.Reader` (for reading)
- `io.Closer` (for closing)

...together forming `io.ReadCloser`.

---

## 15. Closing the Response Body

```go
response, err := http.Get(url)
if err != nil {
    fmt.Println("Error:", err)
    return
}
defer response.Body.Close()
```

This should be done immediately after confirming the request succeeded.

---

## 16. Why Do We Close `response.Body`?

HTTP connections use system/network resources. Closing the body lets Go's HTTP client release resources and, where possible, reuse the underlying connection.

Without properly closing response bodies, a long-running application can eventually run into resource exhaustion. `defer response.Body.Close()` is an essential habit.

---

## 17. Reading the Response Body

```go
import "io"

dataBytes, err := io.ReadAll(response.Body)
if err != nil {
    fmt.Println("Error reading response body:", err)
    return
}
```

`io.ReadAll()` returns `([]byte, error)`.

---

## 18. What Is `[]byte`?

`io.ReadAll()` returns `[]byte` — a slice of raw binary data.

If you know the response is text, convert it to a string:

```go
content := string(dataBytes)
fmt.Println(content)
```

---

## 19. `[]byte` vs `string`

- Use `[]byte` for raw data.
- Use `string(dataBytes)` when treating the bytes as text.

```go
dataBytes, err := io.ReadAll(response.Body)
if err != nil {
    return
}
content := string(dataBytes)
fmt.Println(content)
```

---

## 20. Complete Request Flow

```
URL
 ↓
http.Get()
 ↓
Check error
 ↓
*http.Response
 ↓
Check StatusCode
 ↓
defer response.Body.Close()
 ↓
io.ReadAll(response.Body)
 ↓
[]byte
 ↓
string(dataBytes)
 ↓
Process the response
```

---

## 21. `http.Get()` vs `http.NewRequest()`

`http.Get(url)` is convenient for simple GET requests. But real backend applications often need more control:

- Custom HTTP method
- Headers
- Authentication
- Request body
- Context
- Timeouts
- Cookies

For that, build a request manually:

```go
request, err := http.NewRequest(http.MethodGet, url, nil)

client := &http.Client{}
response, err := client.Do(request)
```

---

## 22. `http.Client`

For more serious applications, use `http.Client`:

```go
client := &http.Client{}
response, err := client.Get(url)
```

This gives more control than the package-level convenience function — you can configure:

- Timeout
- Transport
- Redirect behavior
- Proxies

```go
client := &http.Client{
    Timeout: 10 * time.Second,
}
response, err := client.Get(url)
```

---

## 23. Why Timeouts Matter

Network requests can sometimes hang indefinitely. A backend service should generally avoid waiting forever for another server.

Timeouts are especially important in:

- Microservices
- API integrations
- Payment systems
- Background workers
- Web crawlers
- Production services

---

## 24. JSON APIs

In backend development, you'll often receive JSON instead of HTML:

```json
{
  "id": 1,
  "name": "Abhijeet"
}
```

Instead of treating it as arbitrary text, decode it into a struct:

```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

json.NewDecoder(response.Body).Decode(&user)
```

This is generally preferable to reading the whole body into memory first when processing JSON directly.

---

## 25. HTML vs JSON

- Requesting `https://example.com` → likely HTML response
- Requesting `https://api.example.com/users` → likely JSON response

Inspect `response.Header.Get("Content-Type")` to understand what type of data the server returned.

---

## 26. Important Production Pattern

```go
response, err := client.Do(request)
if err != nil {
    return err
}
defer response.Body.Close()

if response.StatusCode < 200 || response.StatusCode >= 300 {
    // Handle HTTP error
}

body, err := io.ReadAll(response.Body)
if err != nil {
    return err
}
```

This is much closer to what you'll actually use in real backend applications.

---

## 27. Key Takeaways

- `net/http` provides HTTP functionality.
- `http.Get()` makes a GET request.
- `http.Response` represents the server response.
- `response.Status` gives the status text; `response.StatusCode` gives the numeric status.
- `response.Header` contains response headers.
- `response.Body` is a **stream** — always close it.
- `io.ReadAll()` reads the complete body into memory and returns `[]byte`.
- Convert bytes to string with `string(dataBytes)` when appropriate.
- Always check HTTP status codes.
- Use `http.Client` when you need more control (timeouts, transport, etc.).
- For JSON APIs, prefer decoding into structs over raw string reads.
- A successful network request doesn't guarantee a 2xx HTTP response.

---

## 28. Quick Mental Model

```
                HTTP CLIENT
                    │
                    ▼
               http.Get()
                    │
                    ▼
              HTTP REQUEST
                    │
                    ▼
                INTERNET
                    │
                    ▼
                 SERVER
                    │
                    ▼
              HTTP RESPONSE
                    │
          ┌─────────┼─────────┐
          ▼         ▼         ▼
       Status     Headers    Body
          │                   │
          ▼                   ▼
      200 OK              io.Reader
                              │
                              ▼
                         io.ReadAll()
                              │
                              ▼
                            []byte
                              │
                              ▼
                           string
```

---

## 29. What to Learn Next

```
1. HTTP GET
       ↓
2. HTTP POST
       ↓
3. Request Headers
       ↓
4. Request Body
       ↓
5. JSON Encoding / Decoding
       ↓
6. http.NewRequest()
       ↓
7. http.Client
       ↓
8. Context and Cancellation
       ↓
9. Timeouts
       ↓
10. REST API Clients
       ↓
11. HTTP Server
       ↓
12. Routing
       ↓
13. Middleware
```
