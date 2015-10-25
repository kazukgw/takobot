package actiongroups

import (
	act "github.com/kazukgw/takobot/cmd/takobot/actions"
	"github.com/kazukgw/takobot/cmd/takobot/db"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type Nijihan struct {
	db.NewDB
	coa.DoSelf
	act.SendMsg
	db.CloseDB
}

func (ag Nijihan) Schedule() string {
	return "TZ=Asia/Tokyo 0 30 14 * * *"
}

func (ag *Nijihan) Do(ctx coa.Context) error {
	ag.SendMsg.Msg = "ニジ・ハーン"
	ag.SendMsg.Channel = "general"
	return nil
}
