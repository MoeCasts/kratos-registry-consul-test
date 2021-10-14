package main

import (
	"context"
	v1 "kratos-registry-consul-test/api/helloworld/v1"
	"log"
	"testing"
	"time"

	consul "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/hashicorp/consul/api"
)

func TestRegistry(t *testing.T) {
	t.Run("TestConsulGRPCClient", func(t *testing.T) {
		consulConfig := api.DefaultConfig()
		consulConfig.Address = "http://127.0.0.1:8500"
		consulClient, err := api.NewClient(consulConfig)
		if err != nil {
			panic(err)
		}

		endpoint := "discovery:///test-consul-registry-service"
		dis := consul.New(consulClient)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		conn, err := grpc.DialInsecure(ctx, grpc.WithEndpoint(endpoint), grpc.WithDiscovery(dis))
		if err != nil {
			panic(err)
		}

		cancel()
		defer conn.Close()
		gClient := v1.NewGreeterClient(conn)
		callGRPC(gClient)
	})

	t.Run("TestConsulHTTPClient", func(t *testing.T) {
		consulConfig := api.DefaultConfig()
		consulConfig.Address = "http://127.0.0.1:8500"
		consulClient, err := api.NewClient(consulConfig)
		if err != nil {
			panic(err)
		}

		endpoint := "discovery:///test-consul-registry-service"
		dis := consul.New(consulClient)

		hConn, err := http.NewClient(
			context.Background(),
			http.WithMiddleware(
				recovery.Recovery(),
			),
			http.WithEndpoint(endpoint),
			http.WithDiscovery(dis),
		)
		if err != nil {
			log.Fatal(err)
		}
		defer hConn.Close()
		hClient := v1.NewGreeterHTTPClient(hConn)
		callHTTP(hClient)
	})
}

func callGRPC(client v1.GreeterClient) {
	reply, err := client.SayHello(context.Background(), &v1.HelloRequest{Name: "tester"})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[grpc] SayHello: %v\n", reply.GetMessage())
}

func callHTTP(client v1.GreeterHTTPClient) {
	reply, err := client.SayHello(context.Background(), &v1.HelloRequest{Name: "tester"})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[http] SayHello: %v\n", reply.GetMessage())
}
