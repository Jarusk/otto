package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nlopes/slack"
)

func main() {
	api := slack.New(os.Getenv("OTTO_SLACK_TOKEN"))
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(true)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		//fmt.Print("Event Received: ")
		switch ev := msg.Data.(type) {

		case *slack.MessageEvent:
			fmt.Printf("Message: %v\n", ev)

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials\n")
			return

		default:

			// Ignore other events..
		}
	}
}
