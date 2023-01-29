package main

import (
	"fmt"
	"net/http"
	"strings"
)

func defangurl(url string) string {
	return strings.ReplaceAll(url, ".", "[.]")
}

func main() {
	// The shortened URL
	var originalURL string
	fmt.Println("Enter the url:")
	fmt.Scan(&originalURL)
	// Send a GET request to the URL
	resp, err := http.Get(originalURL)
	if err != nil {
		fmt.Println("Error:", err, "Example: https://google.com")
		return
	}
	defer resp.Body.Close()

	// Print the original URL
	fmt.Println("Defanged URL:", defangurl(resp.Request.URL.String()))
}
