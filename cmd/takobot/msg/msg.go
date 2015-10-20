package msg

import (
	"regexp"

	"github.com/kazukgw/takobot/cmd/takobot/models"
	"github.com/kazukgw/takobot/cmd/takobot/store"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/nlopes/slack"
)

var toUserMessagePattern = regexp.MustCompile("^<@([^:]+)>:.*")

type Msg struct {
	*slack.MessageEvent
	ToUser *slack.User
}

func NewMsg(ev *slack.MessageEvent) *Msg {
	msg := &Msg{MessageEvent: ev}
	if toUserMessagePattern.Match([]byte(ev.Text)) {
		res := toUserMessagePattern.FindSubmatch([]byte(ev.Text))
		if len(res) > 0 {
			id := string(res[1])
			u := store.UserByID(id)
			if u.ID != "" {
				msg.ToUser = &u
			}
		}
	}
	return msg
}

func (msg *Msg) ModelsMsg() *models.Msg {
	tmsg := &models.Msg{
		From:     msg.User,
		Channel:  msg.Channel,
		FullBody: msg.Text,
	}
	if msg.ToUser != nil {
		tmsg.To = msg.ToUser.ID
	}
	return tmsg
}
