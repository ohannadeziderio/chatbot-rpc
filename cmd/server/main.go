package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/ohannadeziderio/chatbot-rpc/pkg/chat"
)

func main() {
	chatbotService := new(chat.Chatbot)
  rpc.Register(chatbotService)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Error starting RPC server:", err)
	}
	defer listener.Close()

	log.Println("RPC server started on port 1234")
	rpc.Accept(listener)
}
