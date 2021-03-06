package actions

import (
	ctxs "github.com/kazukgw/takobot/cmd/takobot/contexts"
	"github.com/kazukgw/takobot/cmd/takobot/log"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/nlopes/slack"
)

type GetClient struct {
	client *slack.Client
}

type HasClient interface {
	Client() *slack.Client
}

func (a *GetClient) Do(ctx coa.Context) error {
	log.Action("==> get client")
	mctx := ctx.(*ctxs.MsgContext)
	a.client = mctx.Client
	return nil
}

func (a *GetClient) Client() *slack.Client {
	return a.client
}
