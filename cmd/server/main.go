package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/ohannadeziderio/chatbot-rpc/pkg/chat"
)

type ChatbotService struct {
	chatbot *chat.Chatbot
}

func (s *ChatbotService) SendMessage(msg string, reply *string) error {
	log.Print("Receiving message: ", msg)

	response, err := s.chatbot.SendMessage(msg)
	if err != nil {
		return err
	}
	*reply = response
	return nil
}

func main() {
	chatbot := chat.NewChatbot()
	chatbotService := &ChatbotService{chatbot: chatbot}

	rpc.Register(chatbotService)

	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal("Error starting RPC server:", err)
	}
	defer listener.Close()

	log.Println("RPC server started on port 8090")
	rpc.Accept(listener)
}
