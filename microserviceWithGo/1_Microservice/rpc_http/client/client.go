package client

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/Sungchul-P/go-learning/microserviceWithGo/1_Microservice/rpc_http/contract"
)

const port = 1234

func CreateClient() *rpc.Client {
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return client
}

func PerformRequest(c *rpc.Client) contract.HelloWorldResponse {
	args := &contract.HelloWorldRequest{Name: "World"}
	var reply contract.HelloWorldResponse
	err := c.Call("HelloWorldHandler.HelloWorld", args, &reply)

	if err != nil {
		log.Fatal("error:", err)
	}

	return reply
}
