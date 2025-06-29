// 01_url-status-check.go demonstrates how to check the status of URLs concurrently using goroutines and channels.
// It fetches the status of a list of URLs and prints whether each URL is reachable or not.
package concurrency

import (
	"fmt"
	"net/http"
	"sync"
)

// The checkStatus function will be our "task". It takes a URL,
// makes an HTTP GET request, and prints the status.
func checkStatus(url string, wg *sync.WaitGroup) {
	// defer wg.Done() tells the WaitGroup that this goroutine has finished
	// its work right before the function returns.
	defer wg.Done()

	// Make the HTTP GET request.
	resp, err := http.Get(url)
	if err != nil {
		// If an error occurred (e.g., domain doesn't exist), print the error.
		fmt.Printf("  * [ERROR] Error checking %s: %s\n", url, err.Error())
		return
	}
	// It's important to close the response body to free up resources.
	// We use defer to make sure this happens.
	defer resp.Body.Close()

	// Print the URL and the HTTP status from the response.
	fmt.Printf("  * Status of %s is: %s\n", url, resp.Status)
}

// DemoURLStatusCheck demonstrates how to check the status of multiple URLs concurrently.
func DemoURLStatusCheck() {
	// A list of URLs to check.
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.golang.org",
		"https://www.microsoft.com",
		"http://thisurldoesnotexist.fail", // An example of a failing URL
		"https://www.amazon.com",
		"https://example.com", // A valid URL that should return 200 OK
	}

	// 1. Create a new WaitGroup.
	// A WaitGroup is a counter that blocks the execution of a program
	// until its internal counter becomes 0.
	var wg sync.WaitGroup

	fmt.Println("Starting status checks...")

	// 2. Loop through each URL and start a goroutine to check its status.
	for _, url := range urls {
		// 3. Increment the WaitGroup counter for each URL.
		// This tells the WaitGroup that one more goroutine is starting.
		wg.Add(1)

		// 4. Launch a goroutine.
		// The 'go' keyword starts a new goroutine that executes our
		// checkStatus function concurrently. We pass it the url and
		// a pointer to our WaitGroup.
		go checkStatus(url, &wg) // Start the goroutine.
	}

	// 5. Wait for all goroutines to finish.
	// wg.Wait() blocks the main function until the WaitGroup's counter
	// is zero. The counter is decremented by wg.Done() in each goroutine.
	wg.Wait()

	fmt.Println("All status checks completed.")
}
