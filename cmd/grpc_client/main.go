package main

import (
	"context"
	chatV1 "github.com/globus024/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	address = "localhost:50031"
	chatId  = 1
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to dial server:", err)
	}

	defer conn.Close()
	c := chatV1.NewChatServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.Delete(ctx, &chatV1.DeleteRequest{Id: chatId})
	if err != nil {
		log.Fatal("Failed to create chat:", err)
	}
	log.Printf("Delete chat: %v", chatId)
}
