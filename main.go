package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
    "os"
)

func postRequest(conn net.Conn, req string) {
	payload := strings.NewReader(req)
	httpReq, err := http.NewRequest("POST", os.Getenv("HOST_NAME"), payload)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	httpReq.Header.Set("Content-Type", "text/plain")
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
}

func worker(id int, requests <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := net.Dial("tcp", os.Getenv("HOST_NAME"))
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	for req := range requests {
		postRequest(conn, req)
		fmt.Printf("Worker %d sent request: %s\n", id, req)
	}
}

func main() {
	var wg sync.WaitGroup

	numRequests := os.Getenv("N_REQUESTS")
	numWorkers := os.Getenv("N_WORKERS")

	requests := make(chan string, numRequests)

	// Populate the channel with requests
	for i := 0; i < numRequests; i++ {
		requests <- fmt.Sprintf("Request %d", i)
	}
	close(requests)

	// Launch worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, requests, &wg)
	}

	wg.Wait()
}

