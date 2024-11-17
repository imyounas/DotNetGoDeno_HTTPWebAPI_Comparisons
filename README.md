# .NET 9 JIT vs .NET 9 AOT vs Go 1.23 vs Deno 2.0 WebAPI Req/s Comparison

- This repository contains four projects designed to compare the performance of standard HTTP reqs/sec.
- Each project includes its respective standard libraries for managing HTTP requests and JSON serialization.
- All projects expose a single GET endpoint, /todos, which returns 50 todo items (all with the same structure).
- All projects were tested using Go-Wrk benchmarking tool (https://github.com/tsliwowicz/go-wrk) against 500 concurrent users for 3 minutes.
- All tests were conducted on a local Docker container configured with 2 CPU cores and 1024 MB of resources. 
- Across multiple runs, .NET 9 AOT and .NET 9 JIT, outperformed Deno 2 and Go 1.23 implementations by some margin.

