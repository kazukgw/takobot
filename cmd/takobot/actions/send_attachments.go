package actions

import (
	"github.com/kazukgw/takobot/cmd/takobot/log"
	"github.com/kazukgw/takobot/cmd/takobot/store"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/nlopes/slack"
)

type SendAttachments struct {
	slack.Attachment
	Channel string
}

func (a *SendAttachments) Do(ctx coa.Context) error {
	log.Action("send attachments ==>")
	client := ctx.ActionGroup().(HasClient).Client()
	params := slack.PostMessageParameters{}
	params.Attachments = []slack.Attachment{a.Attachment}
	params.Username = "takobot"
	params.IconURL = "https://encrypted-tbn3.gstatic.com/images?q=tbn:ANd9GcT4Nm9GTh-1aTRmNMdkkAYoCMFHALSj560lxbHA7nYSYjBcptH0JA"
	chanID := store.ChanByName("general").ID
	_, _, err := client.PostMessage(chanID, "", params)
	if err != nil {
		return err
	}
	return nil
}
