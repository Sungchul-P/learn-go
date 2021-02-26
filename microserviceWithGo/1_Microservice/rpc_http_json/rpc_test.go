package main

import (
	"testing"

	"github.com/Sungchul-P/go-learning/microserviceWithGo/1_Microservice/rpc_http_json/client"
	"github.com/Sungchul-P/go-learning/microserviceWithGo/1_Microservice/rpc_http_json/server"
)

func BenchmarkHelloWorldHandlerJSONRPC(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = client.PerformRequest()

	}
}

func init() {
	// start the server
	go server.StartServer()
}
