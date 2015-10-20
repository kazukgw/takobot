package actions

import (
	eh "github.com/kazukgw/takobot/cmd/takobot/errorhandler"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type GetAndSaveMsg struct {
	GetMsgFromCtx
	coa.DoSelf
	SaveMsg

	eh.DoNothing
}

func (ag *GetAndSaveMsg) Do(ctx coa.Context) error {
	ag.SaveMsg.Msg = ag.GetMsgFromCtx.Msg
	return nil
}
