package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"github.com/brianvoe/gofakeit/v7"
	chatV1 "github.com/globus024/chat-server/pkg/chat_v1"
)

const gprcPort = 50031

type chatServer struct {
	chatV1.UnimplementedChatServiceServer // Embedded by value as required
}

func (s *chatServer) SendMessage(ctx context.Context, req *chatV1.SendMessageRequest) (*empty.Empty, error) {
	log.Println("SendMessage method called")
	log.Println("SendMessage method called	with message:", req.GetFrom())
	log.Println("SendMessage method called	with message:", req.Text)

	return &empty.Empty{}, nil
}

func (s *chatServer) Create(ctx context.Context, req *chatV1.CreateRequest) (*chatV1.CreateResponse, error) {
	log.Println("Create method called")
	log.Println("with message:", req.GetUsernames())
	return &chatV1.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *chatServer) Delete(ctx context.Context, req *chatV1.DeleteRequest) (*empty.Empty, error) {
	log.Println("Delete method called")
	log.Println("with message:", req.GetId())
	return &empty.Empty{}, nil
}

func main() {
	grpcServer := grpc.NewServer()
	chatV1.RegisterChatServiceServer(grpcServer, &chatServer{})
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", gprcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("gRPC server listening at :%d", gprcPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
