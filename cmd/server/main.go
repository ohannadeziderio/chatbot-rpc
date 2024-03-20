package server

import (
	"github.com/ohannadeziderio/chatbot-rpc/pkg/chat"
	"log"
	"net"
	"net/rpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8090")

	if err != nil {
		log.Fatalln("Error when init server:", err)

		return
	}

	chatbot := chat.NewChatbot()
	rpc.Register(chatbot)

	log.Println("Server started.")

	rpc.Accept(listener)
}
