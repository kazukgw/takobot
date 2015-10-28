package actiongroups

import (
	act "github.com/kazukgw/takobot/cmd/takobot/actions"
	"github.com/kazukgw/takobot/cmd/takobot/db"
	eh "github.com/kazukgw/takobot/cmd/takobot/errorhandler"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type Nijihan struct {
	db.NewDB
	coa.DoSelf
	act.GetRTMAndSendMsg

	eh.DefaultErrorHandler
}

func (ag Nijihan) Schedule() string {
	return "0 30 05 * * *"
}

func (ag *Nijihan) Do(ctx coa.Context) error {
	ag.SendMsg.Msg = "ニジ・ハーン"
	ag.SendMsg.Channel = "general"
	return nil
}
