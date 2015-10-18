package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/kazukgw/takobot/msghandler"
	"github.com/nlopes/slack"
)

func serveStatus() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "takobot live!!")
	})
	http.ListenAndServe(":"+port, nil)
}

func main() {
	tkn := os.Getenv("SLACK_BOT_TOKEN")
	api := slack.New(tkn)
	api.SetDebug(true)

	go serveStatus()

	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				// Ignore hello

			case *slack.ConnectedEvent:
				fmt.Println("Infos:", ev.Info)
				fmt.Println("Connection counter:", ev.ConnectionCount)
				// Replace #general with your Channel ID
				rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "C0APDGNTF"))

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev)
				fmt.Println(ev.Text)
				outgoing := msghandler.Handle(ev.Text)
				if outgoing != "" {
					rtm.SendMessage(rtm.NewOutgoingMessage(outgoing, "C0APDGNTF"))
				}

			case *slack.PresenceChangeEvent:
				fmt.Printf("Presence Change: %v\n", ev)

			case *slack.LatencyReport:
				fmt.Printf("Current latency: %v\n", ev.Value)

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:

				// Ignore other events..
				// fmt.Printf("Unexpected: %v\n", msg.Data)
			}
		}
	}
}
