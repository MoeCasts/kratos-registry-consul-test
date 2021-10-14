# Quick Start

## Start consul
To start consul with docker:
```
docker run --name consul -d -p 8500:8500 consul
```
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
=== RUN   TestRegistry
=== RUN   TestRegistry/TestConsulGRPCClient
INFO msg=[resolver] update instances: [{"id":"XXX.local","name":"test-consul-registry-service","version":"v0.0.1","metadata":null,"endpoints":["grpc://192.168.2.233:9000","http://192.168.2.233:8000"]}]
2021/10/15 00:21:08 [grpc] SayHello: Hello tester
=== RUN   TestRegistry/TestConsulHTTPClient
2021/10/15 00:21:08 error: code = 503 reason = NODE_NOT_FOUND message = no instances available metadata = map[]
FAIL	command-line-arguments	0.161s
FAIL
```
