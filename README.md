# Comparing HTTP Performance: A Local Benchmark Study of .NET 9, Go, And Deno

- Performance Comparison: This repository contains four projects, each designed to compare the performance of standard HTTP requests per second (req/sec) across different frameworks and implementations.
- Standard Libraries Used: Each project leverages its respective standard libraries for handling HTTP requests and serializing JSON data, ensuring a fair comparison of built-in functionality.
- Common API Endpoint: All projects expose a single GET endpoint, /todos, which returns a fixed set of 50 todo items, each with the same structure, to maintain consistency across tests.
- Benchmarking Tool: Performance tests were conducted using the Go-Wrk benchmarking tool, simulating 500 concurrent users for a duration of 3 minutes to measure HTTP request throughput.
- Test Environment: All tests were run in a local Docker container with 2 CPU cores and 1024 MB of memory, providing an isolated and controlled testing environment.
- Test Results: Across multiple runs, the .NET 9 AOT and .NET 9 JIT implementations outperformed both Deno 2 and Go 1.23 by a some  margin in terms of requests per second.
- However, in terms of low latency, Go has the best results. Due to having more timeouts, it affected its overall average results. Which shows, out of box standard Go http/net library needs more fine tuning for production scenarios.
- For details, you can check this (LinkedIn article)[https://www.linkedin.com/pulse/comparing-http-performance-local-benchmark-study-ofnet-imran-younas-qnvqf]

