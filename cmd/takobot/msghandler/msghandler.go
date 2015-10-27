package msghandler

import (
	"reflect"
	"regexp"

	ags "github.com/kazukgw/takobot/cmd/takobot/actiongroups"
	act "github.com/kazukgw/takobot/cmd/takobot/actions"
	ctxs "github.com/kazukgw/takobot/cmd/takobot/contexts"
	"github.com/kazukgw/takobot/cmd/takobot/db"
	"github.com/kazukgw/takobot/cmd/takobot/log"
	"github.com/kazukgw/takobot/cmd/takobot/msg"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/nlopes/slack"
)

func HandleMsg(
	ev *slack.MessageEvent,
	rtm *slack.RTM,
	client *slack.Client,
) {
	msg := msg.NewMsg(ev)
	mctx := ctxs.NewMsgContext(Routing{}, msg, rtm, client)
	mctx.Exec()
}

type Routing struct {
	db.NewDB
	act.GetAndSaveMsg
	coa.DoSelf
	db.CloseDB
}

func (ag *Routing) Do(ctx coa.Context) error {
	mctx := ctx.(*ctxs.MsgContext)
	msg := ag.GetMsg()
	for ptn, agSource := range commandPatterns {
		if ptn.Match([]byte(msg.Text)) {
			log.ActionGRP("==> ", reflect.TypeOf(agSource))
			newMctx := ctxs.NewMsgContext(agSource, msg, mctx.RTM, mctx.Client)
			return newMctx.Exec()
		}
	}

	log.ActionGRP("==> send registered msg")
	newMctx := ctxs.NewMsgContext(ags.SendRegisteredMsg{}, msg, mctx.RTM, mctx.Client)
	return newMctx.Exec()
}

var commandPatterns = map[*regexp.Regexp]interface{}{
	regexp.MustCompile("pattern[ ]+add:.*"): ags.AddPattern{},
	regexp.MustCompile("pattern[ ]+ls:.*"):  ags.ShowPattern{},
	regexp.MustCompile("pattern[ ]+rm:.*"):  ags.RemovePattern{},
}
