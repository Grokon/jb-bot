package main

import (
	"log"

	"github.com/Grokon/jb-bot/slack"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack/socketmode"
)

func main() {

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	client, err := slack.ConnectToSlackViaSocketmode()
	if err != nil {
		log.Fatal("Failed to connect to Slack via Socket Mode")
	}

	socketmodeHandler := socketmode.NewSocketmodeHandler(client)

	socketmodeHandler.Handle(socketmode.EventTypeConnecting, slack.MiddlewareConnecting)
	socketmodeHandler.Handle(socketmode.EventTypeConnectionError, slack.MiddlewareConnectionError)
	socketmodeHandler.Handle(socketmode.EventTypeConnected, slack.MiddlewareConnected)

	// //\\ EventTypeEventsAPI //\\
	// // Handle all EventsAPI
	// SocketmodeHandler.Handle(socketmode.EventTypeEventsAPI, middlewareEventsAPI)

	// SocketmodeHandler.Handle(socketmode.EventTypeHello, func(evt *socketmode.Event, cli *socketmode.Client) {
	// 	log.Println("Hello received")
	// })

	// Build a Slack App Home in Golang Using Socket Mode

	// run socketmode
	socketerror := socketmodeHandler.RunEventLoop()
	if socketerror != nil {
		log.Fatal(socketerror)
	}

}
