# Quick Start

## Start service
```
kratos run
```

## Run tests

```
go test ./cmd/kratos-registry-consul-test/main_test.go -v
```

output:

```
--- FAIL: TestHTTPClient (0.00s)
    --- FAIL: TestHTTPClient/TestHTTPClientWithTime (0.00s)
panic: error: code = 500 reason =  message = parsing field "time": parsing time "" as "2006-01-02T15:04:05.999999999Z07:00": cannot parse "" as "2006" metadata = map[] [recovered]
```
