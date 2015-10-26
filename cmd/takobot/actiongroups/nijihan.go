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
	act.SendMsg
	db.CloseDB
}

func (ag Nijihan) Schedule() string {
	return "0 0 13 * * *"
}

func (ag *Nijihan) Do(ctx coa.Context) error {
	fmt.Println("do action group: Nijihan")
	ag.SendMsg.Msg = "ニジ・ハーン"
	ag.SendMsg.Channel = "general"
	return nil
}

func (ag *Nijihan) HandleError(ctx coa.Context, err error) error {
	fmt.Println(err.Error())
	return err
}
