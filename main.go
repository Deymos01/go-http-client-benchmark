package main

import (
	"flag"
	"fmt"
	"github.com/Deymos01/go-http-client-benchmark/benchmark"
	"github.com/Deymos01/go-http-client-benchmark/clients"
	"os"
)

func main() {
	var (
		requestCount  int
		url           string
		benchmarkMode string // "seq" for sequential, "par" for parallel
		clientName    string
	)

	flag.IntVar(&requestCount, "n", 1000, "Number of requests to send")
	flag.StringVar(&url, "url", "http://localhost:8080/ping", "Target URL")
	flag.StringVar(&benchmarkMode, "mode", "seq", "Benchmark mode: 'seq' for sequential, 'par' for parallel")
	flag.StringVar(&clientName, "client", "nethttp", "HTTP client to use")
	flag.Parse()

	var cl clients.Client
	switch clientName {
	case "nethttp":
		cl = clients.NewHTTPClient(url)
	case "fasthttp":
		cl = clients.NewFastHttpClient(url)
	default:
		fmt.Println("Unknown client. Use defaults: 'nethttp' or 'fasthttp'.")
		os.Exit(1)
	}

	switch benchmarkMode {
	case "seq":
		benchmark.RunBenchmarkSequential(fmt.Sprintf("Bench client %s (Sequential mode)", clientName), cl, requestCount)
	case "par":
		benchmark.RunBenchmarkParallel(fmt.Sprintf("Bench client %s (Parallel mode)", clientName), cl, requestCount)
	default:
		fmt.Println("Unknown benchmark mode. Use 'seq' for sequential or 'par' for parallel.")
	}
}
