package contexts

import (
	"reflect"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/nlopes/slack"
)

type RTMContext struct {
	actionGroup coa.ActionGroup
	rtm         *slack.RTM
}

func NewRTMContext(agSource interface{}, rtm *slack.RTM) *RTMContext {
	agType := reflect.TypeOf(agSource)
	if ag, ok := reflect.New(agType).Interface().(coa.ActionGroup); ok {
		return &RTMContext{ag, rtm}
	}
	return nil
}

func (ctx *RTMContext) ActionGroup() coa.ActionGroup {
	return ctx.actionGroup
}

func (ctx *RTMContext) RTM() *slack.RTM {
	return ctx.rtm
}

func (ctx *RTMContext) Exec() error {
	return coa.Exec(ctx.ActionGroup(), ctx)
}
