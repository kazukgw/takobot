package actions

import (
	"github.com/kazukgw/takobot/cmd/takobot/store"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/nlopes/slack"
)

type SendAttachments struct {
	slack.Attachment
	Channel string
}

func (a *SendAttachments) Do(ctx coa.Context) error {
	client := ctx.ActionGroup().(HasClient).Client()
	params := slack.PostMessageParameters{}
	params.Attachments = []slack.Attachment{a.Attachment}
	params.Username = "takobot"
	chanID := store.ChanByName("general").ID
	_, _, err := client.PostMessage(chanID, "", params)
	if err != nil {
		return err
	}
	return nil
}
