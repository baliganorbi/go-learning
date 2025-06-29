// 02_worker-pool.go demonstrates a worker pool pattern in Go.
// Adapt the URL Status Checker. Instead of one goroutine per URL,
// create a fixed pool of workers that pull URLs from a "jobs" channel
// and process them.
package concurrency

import (
	"fmt"
	"net/http"
	"time"
)

// The worker function is now our long-lived goroutine.
// It receives jobs from the 'jobs' channel and sends results to the 'results' channel.
// We use directional channels (jobs <-chan, results chan<-) for type safety.
func worker(id int, jobs <-chan string, results chan<- string) {
	// This 'for range' loop will run until the 'jobs' channel is closed
	// and all values have been received from it.
	for url := range jobs {
		fmt.Printf("  * Worker %d starting job for URL: %s\n", id, url)

		resp, err := http.Get(url)
		var result string
		if err != nil {
			result = fmt.Sprintf("    - [ERROR] Error checking %s: %s", url, err.Error())
		} else {
			resp.Body.Close()
			result = fmt.Sprintf("    - Status of %s is: %s", url, resp.Status)
		}

		// Send the result string to the results channel.
		results <- result
		time.Sleep(1000 * time.Millisecond) // To better visualize workers picking jobs
	}
}

// DemoWorkerPool demonstrates a worker pool pattern for checking URL statuses.
func DemoWorkerPool() {
	// A list of URLs to check.
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.golang.org",
		"https://www.microsoft.com",
		"http://thisurldoesnotexist.fail", // An example of a failing URL
		"https://www.amazon.com",
		"https://example.com", // A valid URL that should return 200 OK
		"https://www.wikipedia.org",
		"https://www.stackoverflow.com",
		"https://www.reddit.com",
		"https://www.twitter.com",
		"https://www.facebook.com",
		"https://www.linkedin.com",
		"https://www.youtube.com",
		"https://www.instagram.com",
		"https://www.netflix.com",
	}

	numJobs := len(urls)

	// 1. Create our two channels: one for jobs and one for results.
	// These channels are buffered to allow multiple jobs to be sent
	jobs := make(chan string, numJobs)    // Channel for jobs (URLs)
	results := make(chan string, numJobs) // Channel for results

	// 2. Start a fixed number of workers.
	// We'll start 3 workers. They will initially be blocked because
	// the 'jobs' channel is empty.
	const numWorkers = 3
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// 3. Send all the jobs to the 'jobs' channel.
	for _, url := range urls {
		jobs <- url
	}

	// 4. Close the 'jobs' channel.
	// This is a CRITICAL step. It signals to the workers (in their 'for range' loop)
	// that no more jobs will be sent. Once they finish their current jobs, they will exit.
	close(jobs)

	fmt.Println("Starting status checks with worker pool...")

	// 5. Collect all the results.
	// We'll loop 'numJobs' times to ensure we get a result for every job we sent.
	fmt.Println("\n--- Results ---")
	for a := 1; a <= numJobs; a++ {
		// This will block until a result is received.
		result := <-results
		fmt.Println(result)
	}

	fmt.Println("\nAll jobs complete.")
}
