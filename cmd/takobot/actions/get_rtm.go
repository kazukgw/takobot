package actions

import (
	ctxs "github.com/kazukgw/takobot/cmd/takobot/contexts"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/nlopes/slack"
)

type GetRTM struct {
	rtm *slack.RTM
}

type HasRTM interface {
	RTM() *slack.RTM
}

func (a *GetRTM) Do(ctx coa.Context) error {
	mctx := ctx.(*ctxs.MsgContext)
	a.rtm = mctx.RTM
	return nil
}

func (a *GetRTM) RTM() *slack.RTM {
	return a.rtm
}
