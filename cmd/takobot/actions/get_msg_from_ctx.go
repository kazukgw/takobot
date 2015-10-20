package actions

import (
	ctxs "github.com/kazukgw/takobot/cmd/takobot/contexts"
	"github.com/kazukgw/takobot/cmd/takobot/msg"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type GetMsgFromCtx struct {
	*msg.Msg
}

type HasMsg interface {
	GetMsg() *msg.Msg
}

func (a *GetMsgFromCtx) Do(ctx coa.Context) error {
	mctx := ctx.(*ctxs.MsgContext)
	a.Msg = mctx.Msg
	return nil
}

func (a *GetMsgFromCtx) GetMsg() *msg.Msg {
	return a.Msg
}
