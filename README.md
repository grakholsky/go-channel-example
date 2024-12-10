# Go Channels Example
Test app

## Building

    go build .

## Testing

Install [vegeta](https://github.com/tsenart/vegeta) HTTP load testing tool and library.

```
$ go get -u github.com/tsenart/vegeta
```

## Benchmark log
- CPU: Core i7-8565U
- Memory: 8 GB
```
$ echo "GET http://localhost:8080/request" | vegeta attack -rate=10000 -duration=60s | vegeta report
Requests      [total, rate]            600000, 9999.92
Duration      [total, attack, wait]    1m0.0004895s, 1m0.0004895s, 0s
Latencies     [mean, 50, 95, 99, max]  33.15µs, 0s, 0s, 993.62µs, 18.9503ms
Bytes In      [total, mean]            1200000, 2.00
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:600000
Error Set:
```
