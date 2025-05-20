package benchmark

import (
	"fmt"
	"github.com/Deymos01/go-http-client-benchmark/clients"
	"sync"
	"time"
)

func RunBenchmarkSequential(name string, client clients.Client, reqCount int) {
	start := time.Now()

	var failed int

	for i := 0; i < reqCount; i++ {
		err := client.MakeGetRequest()
		if err != nil {
			failed++
			fmt.Printf("[%s] Request %d failed: %v\n", name, i, err)
		}
	}

	duration := time.Since(start)
	success := reqCount - failed
	avgTime := duration / time.Duration(reqCount)

	fmt.Printf("[%s] Completed %d requests in %v (Failed: %d, Avg Time: %v)\n", name, success, duration, failed, avgTime)
	fmt.Printf("Success Rate: %.2f%%\n", float64(success)/float64(reqCount)*100)
}

func RunBenchmarkParallel(name string, client clients.Client, reqCount int) {
	start := time.Now()

	var wg sync.WaitGroup
	var failed int
	mu := sync.Mutex{}

	for i := 0; i < reqCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := client.MakeGetRequest()
			if err != nil {
				mu.Lock()
				failed++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	duration := time.Since(start)
	success := reqCount - failed
	avgTime := duration / time.Duration(reqCount)

	fmt.Printf("[%s] Completed %d requests in %v (Failed: %d, Avg Time: %v)\n", name, success, duration, failed, avgTime)
	fmt.Printf("Success Rate: %.2f%%\n", float64(success)/float64(reqCount)*100)
}
