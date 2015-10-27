package actions

import (
	eh "github.com/kazukgw/takobot/cmd/takobot/errorhandler"
	"github.com/kazukgw/takobot/cmd/takobot/log"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type GetRTMAndSendMsg struct {
	GetRTM
	SendMsg

	eh.DoNothing
}

func (ag *GetRTMAndSendMsg) Do(ctx coa.Context) error {
	log.Action("==> get rtm and send msg")
	return nil
}
