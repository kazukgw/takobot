package errorhandler

import (
	"errors"

	ctxs "github.com/kazukgw/takobot/cmd/takobot/contexts"
	"github.com/kazukgw/takobot/cmd/takobot/db"
	"github.com/kazukgw/takobot/cmd/takobot/log"
	"github.com/kazukgw/takobot/cmd/takobot/store"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

var ErrNonError = errors.New("")

type DoNothing struct {
}

func (eh *DoNothing) HandleError(ctx coa.Context, err error) error {
	return err
}

type DefaultErrorHandler struct {
}

func (eh *DefaultErrorHandler) HandleError(ctx coa.Context, err error) error {
	ag := ctx.ActionGroup()
	if hasdb, ok := ag.(db.HasDB); ok {
		hasdb.DB().Close()
	}
	log.Error(err)
	return err
}

type WakaranErrHandler struct {
}

func (eh *WakaranErrHandler) HandleError(ctx coa.Context, err error) error {
	if err == ErrNonError {
		return err
	}
	err = (&DefaultErrorHandler{}).HandleError(ctx, err)
	mctx := ctx.(*ctxs.MsgContext)
	chanName := mctx.Msg.Channel
	rtm := mctx.RTM
	rtm.SendMessage(rtm.NewOutgoingMessage("すまん。わからんわ", store.ChanByName(chanName).ID))
	return err
}
