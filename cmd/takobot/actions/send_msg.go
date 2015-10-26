package actions

import (
	"fmt"

	"github.com/kazukgw/takobot/cmd/takobot/store"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type SendMsg struct {
	Msg     string
	Channel string
}

func (a *SendMsg) Do(ctx coa.Context) error {
	rtm := ctx.ActionGroup().(HasRTM).RTM()
	fmt.Printf("send msg to channel: %v msg: %v\n", a.Channel, a.Msg)
	rtm.SendMessage(rtm.NewOutgoingMessage(a.Msg, store.ChanByName(a.Channel).ID))
	return nil
}
