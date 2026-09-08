package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("URL handling")

	urlString := "https://projectlog.hashnode.dev/audio-subtitle-pipeline-rebuild"

	// Parse the URL.
	result, err := url.Parse(urlString)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	// Inspect different parts of the URL.
	fmt.Println("Parsed URL:", result)
	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("Raw Query:", result.RawQuery)
	fmt.Println("Fragment:", result.Fragment)

	// Get query parameters.
	qparams := result.Query()

	fmt.Printf("Query Parameters: %T\n", qparams)

	for key, values := range qparams {
		fmt.Println("Key:", key, "Values:", values)
	}

	// Construct a URL.
	partsOfURL := &url.URL{
		Scheme:   "https",
		Host:     "projectlog.hashnode.dev",
		Path:     "/audio-subtitle-pipeline-rebuild",
		RawQuery: "param1=value1&param2=value2",
	}

	mainURL := partsOfURL.String()

	fmt.Println("Constructed URL:", mainURL)
}

// HTTP example:
//
// func makeRequest() {
// 	urlString := "https://example.com"
//
// 	response, err := http.Get(urlString)
// 	if err != nil {
// 		fmt.Println("Error making request:", err)
// 		return
// 	}
// 	defer response.Body.Close()
//
// 	fmt.Println("Status:", response.Status)
// 	fmt.Println("Status Code:", response.StatusCode)
//
// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response body:", err)
// 		return
// 	}
//
// 	fmt.Println(string(body))
// }

// A safer way to construct query parameters:
//
// func buildURL() {
// 	values := url.Values{}
//
// 	values.Set("name", "abhijeet")
// 	values.Set("page", "2")
//
// 	u := &url.URL{
// 		Scheme:   "https",
// 		Host:     "example.com",
// 		Path:     "/users",
// 		RawQuery: values.Encode(),
// 	}
//
// 	fmt.Println(u.String())
// }
