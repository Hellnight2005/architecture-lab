// Go HTTP Requests — Code

package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://profitable-azure.vercel.app/"

func main() {
	fmt.Println("Web Request")

	// Send an HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Always close the response body once we're done with it
	defer response.Body.Close()

	// Inspect the response
	fmt.Printf("Response type: %T\n", response)
	fmt.Println("Status:", response.Status)
	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content Length:", response.ContentLength)

	// Read the response body
	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Convert []byte to string
	content := string(dataBytes)

	fmt.Println("Response Body:")
	fmt.Println(content)
}

//Extended version — with status code check + client timeout

// package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"time"
// )

// const url = "https://profitable-azure.vercel.app/"

// func main() {
// 	fmt.Println("Web Request")

// 	// Use http.Client for more control (timeout, transport, etc.)
// 	client := &http.Client{
// 		Timeout: 10 * time.Second,
// 	}

// 	response, err := client.Get(url)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer response.Body.Close()

// 	fmt.Println("Status:", response.Status)
// 	fmt.Println("Status Code:", response.StatusCode)
// 	fmt.Println("Content Length:", response.ContentLength)

// 	// Check for a successful HTTP status (2xx range)
// 	if response.StatusCode < 200 || response.StatusCode >= 300 {
// 		fmt.Println("Non-success status code received:", response.StatusCode)
// 		return
// 	}

// 	// Inspect Content-Type before deciding how to handle the body
// 	contentType := response.Header.Get("Content-Type")
// 	fmt.Println("Content-Type:", contentType)

// 	dataBytes, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response body:", err)
// 		return
// 	}

// 	content := string(dataBytes)

// 	fmt.Println("Response Body:")
// 	fmt.Println(content)
// }
