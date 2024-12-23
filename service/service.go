// Service implementation (what you implement)
package service

import (
	"context"
	"fmt"
	"log"

	"connectrpc.com/connect"
	v1 "github.com/naoyafurudono/proto-cli/gen/greet/v1" // generated by protoc-gen-go
	"github.com/naoyafurudono/proto-cli/gen/greet/v1/greetv1connect"
)

type GreetServer struct{}

var _ greetv1connect.GreetServiceHandler = &GreetServer{}

func (s *GreetServer) Hello(
	ctx context.Context,
	req *connect.Request[v1.HelloRequest],
) (*connect.Response[v1.HelloResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&v1.HelloResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.GetName()),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func (s *GreetServer) Thanks(
	ctx context.Context,
	req *connect.Request[v1.ThanksRequest],
) (*connect.Response[v1.ThanksResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&v1.ThanksResponse{
		Greeting: fmt.Sprintf("Thanks %s for your %s!", req.Msg.GetName(), req.Msg.GetWhy()),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}
