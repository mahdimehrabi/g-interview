package main

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	rand2 "math/rand"
	"net/http"
	"time"
)

type Req struct {
	Message string `json:"Message"`
}

func generateRandomBytes(n int) ([]byte, error) {
	randomBytes := make([]byte, 50+n)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	return randomBytes, nil
}

func main() {
	s := rand2.NewSource(time.Now().UnixNano())
	r := rand2.New(s)
	// Seed the random number generator with the current time in nanoseconds
	start := time.Now()

	// Define the URL to send the POST request to
	url := "http://localhost:8000/api/message/send"

	// Get the number of requests to send from command line input
	numRequests := flag.Int("n", 1, "Number of requests to send")
	flag.Parse()

	// Counters for successful and failed requests
	successfulRequests := 0
	failedRequests := 0

	for i := 0; i < *numRequests; i++ {
		// Generate random data for the Req
		b, err := generateRandomBytes(r.Intn(8192 - 50))
		if err != nil {
			log.Fatal(err)
		}
		trade := Req{
			Message: string(b),
		}

		// Convert the Req to a JSON payload
		payload, err := json.Marshal(trade)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			failedRequests++
			continue
		}

		// Send the POST request
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
		if err != nil {
			fmt.Println("Error sending POST request:", err)
			failedRequests++
			continue
		}
		resp.Body.Close()

		// Read and parse the response
		if resp.StatusCode == http.StatusOK {
			fmt.Printf("Produced item %d successfully\n", i)
			successfulRequests++
		} else {
			fmt.Printf("Received a non-OK response: %d for item:%d\n", resp.StatusCode, i)
			failedRequests++
		}
		time.Sleep(420 * time.Nanosecond)
	}
	end := time.Now()
	fmt.Printf("Successful requests: %d\n", successfulRequests)
	fmt.Printf("Failed requests: %d\n", failedRequests)
	fmt.Printf("Benchmark: %s \n", end.Sub(start))
}
