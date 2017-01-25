package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sromku/go-gitter"
)

func main() {

	// Setup
	gitter_token := ""
	client := gitter.New(gitter_token)

	var CURRENT_ROOM_ID string = ""

	// Execution
	receiver := client.Stream(CURRENT_ROOM_ID)
	go client.Listen(receiver)

	sender := make(chan string, 10)
	go get_input(sender)

	// Interaction
	for {
		select {
		case msg := <-receiver.Event:
			switch ev := msg.Data.(type) {
			case *gitter.MessageReceived:
				fmt.Printf("[%s]: %s", ev.Message.From.Username, ev.Message.Text)
			case *gitter.GitterConnectionClosed:
				fmt.Printf("!!Gitter Connection Closed!!")
				panic("!!Gitter Connection Closed!!")
			}

		case send := <-sender:
			client.SendMessage(CURRENT_ROOM_ID, send)
		}
	}
}

func get_input(receiver chan string) {
	// Get user input
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		receiver <- input
	}
}
