package actions

import (
	eh "github.com/kazukgw/takobot/cmd/takobot/errorhandler"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type SendMsgWithPattern struct {
	MsgWithPattern
	coa.DoSelf
	SendMsg

	eh.DoNothing
}

func (ag *SendMsgWithPattern) Do(ctx coa.Context) error {
	ag.SendMsg.Msg = ag.MsgWithPattern.ResultMsg
	ag.SendMsg.Channel = "general"
	return nil
}
