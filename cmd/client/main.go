package main

import (
	"bufio"
	"log"
	"net/rpc"
	"os"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Error connecting to RPC server:", err)
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		log.Print("You: ")
		message, _ := reader.ReadString('\n')

		var response string
		err = client.Call("ChatbotService.SendMessage", message, &response)
		if err != nil {
			log.Fatal("Error calling RPC method: ", err)
		}

		log.Println("Chatbot: ", response)
	}
}
