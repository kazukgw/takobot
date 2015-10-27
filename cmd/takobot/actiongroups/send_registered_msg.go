package actiongroups

import (
	act "github.com/kazukgw/takobot/cmd/takobot/actions"
	"github.com/kazukgw/takobot/cmd/takobot/db"
	eh "github.com/kazukgw/takobot/cmd/takobot/errorhandler"
	"github.com/kazukgw/takobot/cmd/takobot/models"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type SendRegisteredMsg struct {
	db.NewDB
	act.GetMsgFromCtx
	act.Filtering
	coa.DoSelf
	act.SendMsgWithPattern
	db.CloseDB

	eh.WakaranErrHandler
}

func (ag *SendRegisteredMsg) PreExec(ctx coa.Context) error {
	ag.Filtering.ToPermitUser = []string{"takobot"}
	return nil
}

func (ag *SendRegisteredMsg) Do(ctx coa.Context) error {
	ag.SendMsgWithPattern.Patterns = models.CompiledPatterns
	ag.SendMsgWithPattern.Source = ag.GetMsgFromCtx.Msg.Text
	return nil
}
