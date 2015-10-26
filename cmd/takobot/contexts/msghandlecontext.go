package contexts

import (
	"reflect"

	"github.com/kazukgw/takobot/cmd/takobot/msg"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/nlopes/slack"
)

type MsgContext struct {
	actionGroup coa.ActionGroup
	*msg.Msg
	*slack.RTM
	*slack.Client
}

func NewMsgContext(
	agSource interface{},
	m *msg.Msg,
	rtm *slack.RTM,
	client *slack.Client,
) *MsgContext {
	agType := reflect.TypeOf(agSource)
	if ag, ok := reflect.New(agType).Interface().(coa.ActionGroup); ok {
		return &MsgContext{ag, m, rtm, client}
	}
	return nil
}

func (ctx *MsgContext) ActionGroup() coa.ActionGroup {
	return ctx.actionGroup
}

func (ctx *MsgContext) Exec() error {
	return coa.Exec(ctx.ActionGroup(), coa.Context(ctx))
}
