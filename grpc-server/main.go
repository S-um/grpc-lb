package main

import (
	"context"
	"fmt"
	mygrpc "github.com/s-um/grpc-lb/pkg/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var cnt = 0

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	mygrpc.RegisterUserServer(grpcServer, NewServer())
	if err := grpcServer.Serve(lis); err != nil {
	}
}

type Server struct {
	mygrpc.UnimplementedUserServer
	name string
}

func NewServer() *Server {
	return &Server{UnimplementedUserServer: mygrpc.UnimplementedUserServer{}, name: GetHostName()}
}

func (s *Server) HelloWorld(context.Context, *mygrpc.HelloWorldRequest) (*mygrpc.HelloWorldResponse, error) {
	fmt.Println(cnt)
	cnt++
	resp := mygrpc.HelloWorldResponse{Name: s.name}
	return &resp, nil
}

func GetHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "default"
	}
	return hostname
}
