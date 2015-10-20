package actions

import (
	"fmt"

	ctxs "github.com/kazukgw/takobot/cmd/takobot/contexts"
	"github.com/kazukgw/takobot/cmd/takobot/store"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type SendMsg struct {
	Msg     string
	Channel string
}

func (a *SendMsg) Do(ctx coa.Context) error {
	mctx := ctx.(*ctxs.MsgContext)
	rtm := mctx.RTM
	fmt.Printf("send msg to channel: %v msg: %v\n", a.Channel, a.Msg)
	rtm.SendMessage(rtm.NewOutgoingMessage(a.Msg, store.ChanByName(a.Channel).ID))
	return nil
}
