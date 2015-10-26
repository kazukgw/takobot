package actiongroups

import (
	"fmt"

	act "github.com/kazukgw/takobot/cmd/takobot/actions"
	"github.com/kazukgw/takobot/cmd/takobot/db"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type Nijihan struct {
	db.NewDB
	coa.DoSelf
	act.GetRTMAndSendMsg
	db.CloseDB
}

func (ag Nijihan) Schedule() string {
	return "0 30 03 * * *"
}

func (ag *Nijihan) Do(ctx coa.Context) error {
	ag.SendMsg.Msg = "ニジ・ハーン"
	ag.SendMsg.Channel = "general"
	return nil
}

func (ag *Nijihan) HandleError(ctx coa.Context, err error) error {
	fmt.Println(err.Error())
	return err
}
