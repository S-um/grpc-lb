package main

import (
	"context"
	"fmt"
	mygrpc "github.com/s-um/grpc-lb/pkg/grpc"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	addr = GetString("ADDR", "localhost:8080")
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := mygrpc.NewUserClient(conn)
	_duration, err := strconv.ParseInt(GetString("DURATION", "1"), 10, 64)
	duration := time.Duration(_duration)
	if err != nil {
		log.Fatalf("wrong env DURATION value set : %v", err)
	}
	for i := 0;; i++ {
		resp, err := client.HelloWorld(context.TODO(), &mygrpc.HelloWorldRequest{})
		if err != nil {
			fmt.Printf("Err : [%v]\n", err)
		} else {
			fmt.Println(i, ":", resp.Name)
		}
		time.Sleep(duration * time.Second)
	}
}

func GetString(key, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return val
}
