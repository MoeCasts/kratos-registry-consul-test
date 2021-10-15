package main

import (
	"context"
	v1 "kratos-registry-consul-test/api/helloworld/v1"
	"log"
	"testing"

	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestHTTPClient(t *testing.T) {
	t.Run("TestHTTPClientWithTime", func(*testing.T) {
		conn, err := http.NewClient(
			context.Background(),
			http.WithEndpoint("127.0.0.1:8000"),
		)
		if err != nil {
			panic(err)
		}

		client := v1.NewGreeterHTTPClient(conn)
		reply, err := client.SayHello(context.Background(), &v1.HelloRequest{
			Name: "tester",
			Time: timestamppb.Now(),
		})

		if err != nil {
			panic(err)
		}

		log.Printf("reply= %v", reply)
	})
}
