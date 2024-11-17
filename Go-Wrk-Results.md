#### Go Web API
```
Running 180s test @ http://localhost:6565/todos
  500 goroutine(s) running concurrently
1516178 requests in 2m59.694421177s, 8.18GB read
Requests/sec:           8437.54
Transfer/sec:           46.62MB
Overall Requests/sec:   8381.44
Overall Transfer/sec:   46.31MB
Fastest Request:        0s
Avg Req Time:           59.258ms
Slowest Request:        1.036927s
Number of Errors:       150
Error Counts:           net/http: timeout awaiting response headers=150
10%:                    756µs
50%:                    1.609ms
75%:                    1.945ms
99%:                    2.27ms
99.9%:                  2.281ms
99.9999%:               2.282ms
99.99999%:              2.282ms
stddev:                 50.762ms
```

#### .NET 9 AOT  Web API
```
Running 180s test @ http://localhost:8585/todos
  500 goroutine(s) running concurrently
1875733 requests in 2m59.707881678s, 8.44GB read
Requests/sec:           10437.68
Transfer/sec:           48.07MB
Overall Requests/sec:   10367.64
Overall Transfer/sec:   47.75MB
Fastest Request:        0s
Avg Req Time:           47.903ms
Slowest Request:        1.043039s
Number of Errors:       129
Error Counts:           net/http: timeout awaiting response headers=129
10%:                    2.728ms
50%:                    14.15ms
75%:                    16.5ms
99%:                    17.926ms
99.9%:                  17.972ms
99.9999%:               17.979ms
99.99999%:              17.979ms
stddev:                 17.879ms
```

#### .NET 9 JIT  Web API
```
Running 180s test @ http://localhost:7575/api/todos
  500 goroutine(s) running concurrently
1795062 requests in 2m59.553283835s, 8.07GB read
Requests/sec:           9997.38
Transfer/sec:           46.04MB
Overall Requests/sec:   9917.73
Overall Transfer/sec:   45.67MB
Fastest Request:        507µs
Avg Req Time:           50.013ms
Slowest Request:        1.031967s
Number of Errors:       210
Error Counts:           net/http: timeout awaiting response headers=210
10%:                    6.053ms
50%:                    17.22ms
75%:                    19.148ms
99%:                    20.558ms
99.9%:                  20.614ms
99.9999%:               20.624ms
99.99999%:              20.624ms
stddev:                 18.204ms
```
#### DENO 2.0 Web API
```
Running 180s test @ http://localhost:5555/todos
  500 goroutine(s) running concurrently
1394966 requests in 2m59.868721178s, 5.09GB read
Requests/sec:           7755.47
Transfer/sec:           28.99MB
Overall Requests/sec:   7711.04
Overall Transfer/sec:   28.83MB
Fastest Request:        165µs
Avg Req Time:           64.47ms
Slowest Request:        1.032383s
Number of Errors:       68
Error Counts:           net/http: timeout awaiting response headers=68
10%:                    1.628ms
50%:                    26.926ms
75%:                    29.376ms
99%:                    31.088ms
99.9%:                  31.136ms
99.9999%:               31.142ms
99.99999%:              31.142ms
stddev:                 15.607ms
```