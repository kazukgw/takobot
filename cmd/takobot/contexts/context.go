package contexts

import (
	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type Context struct {
	actionGroup coa.ActionGroup
}

func NewContext(ag coa.ActionGroup) *Context {
	return &Context{ag}
}

func (ctx *Context) ActionGroup() coa.ActionGroup {
	return ctx.actionGroup
}

func (ctx *Context) Exec() error {
	return coa.Exec(ctx.ActionGroup(), ctx)
}
