package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	hellopb "main/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type helloServer struct {
	hellopb.UnimplementedGreetingServiceServer
}

func NewServer() *helloServer {
	return &helloServer{}
}

func (s *helloServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	// リクエストからnameフィールドを取り出して
	// "Hello, [名前]!"というレスポンスを返す
	return &hellopb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}


func main() {
	
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()

	hellopb.RegisterGreetingServiceServer(s, NewServer())
	reflection.Register(s)

	go func () {
		log.Printf("start GRPC server port: %v", port)
		s.Serve(listener)
	}()


	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stop gracefully")
	s.GracefulStop()
}