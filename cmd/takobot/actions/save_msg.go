package actions

import (
	"github.com/kazukgw/takobot/cmd/takobot/db"
	"github.com/kazukgw/takobot/cmd/takobot/log"
	"github.com/kazukgw/takobot/cmd/takobot/msg"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type SaveMsg struct {
	*msg.Msg
}

func (a *SaveMsg) Do(ctx coa.Context) error {
	log.Action("save msg ==>")
	if a.Msg == nil {
		return nil
	}

	ag := ctx.ActionGroup()
	ag.(db.HasDB).DB().Create(a.Msg.ModelsMsg())
	return nil
}
