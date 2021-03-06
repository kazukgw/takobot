package main

import (
	"fmt"
	"os"

	ags "github.com/kazukgw/takobot/cmd/takobot/actiongroups"
	ctxs "github.com/kazukgw/takobot/cmd/takobot/contexts"
	"github.com/kazukgw/takobot/cmd/takobot/cron"
	mh "github.com/kazukgw/takobot/cmd/takobot/msghandler"
	"github.com/kazukgw/takobot/cmd/takobot/store"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/nlopes/slack"
)

func HandleEvent() {
	tkn := os.Getenv("SLACK_BOT_TOKEN")
	api := slack.New(tkn)
	// api.SetDebug(true)
	ctxs.NewContext(&ags.LoadPattern{}).Exec()

	rtm := api.NewRTM()
	go rtm.ManageConnection()
	go cron.Init(rtm, api)

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			// fmt.Print("Event Received: ")
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				// Ignore hello

			case *slack.ConnectedEvent:
				// fmt.Println("Infos:", ev.Info)
				users, _ := api.GetUsers()
				store.MakeUserMap(users)
				channels, _ := api.GetChannels(true)
				store.MakeChanMap(channels)
				msg := rtm.NewOutgoingMessage(
					"あーえー気分やわ",
					store.ChanByName("general").ID,
				)
				rtm.SendMessage(msg)

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev)
				go mh.HandleMsg(ev, rtm, api)

			case *slack.PresenceChangeEvent:
				fmt.Printf("Presence Change: %v\n", ev)

			case *slack.LatencyReport:
				// fmt.Printf("Current latency: %v\n", ev.Value)

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
