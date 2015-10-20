package errorhandler

import (
	"errors"
	"fmt"

	ctxs "github.com/kazukgw/takobot/cmd/takobot/contexts"
	"github.com/kazukgw/takobot/cmd/takobot/store"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

var ErrNonError = errors.New("")

type DoNothing struct {
}

func (eh *DoNothing) HandleError(ctx coa.Context, err error) error {
	return err
}

type WakaranErrHandler struct {
}

func (eh *WakaranErrHandler) HandleError(ctx coa.Context, err error) error {
	if err == ErrNonError {
		return err
	}

	mctx := ctx.(*ctxs.MsgContext)
	fmt.Println(err.Error())
	chanName := mctx.Msg.Channel
	rtm := mctx.RTM
	rtm.SendMessage(rtm.NewOutgoingMessage("すまん。わからんわ", store.ChanByName(chanName).ID))
	return err
}
