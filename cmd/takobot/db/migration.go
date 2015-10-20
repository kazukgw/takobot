package db

import (
	ms "github.com/kazukgw/takobot/cmd/takobot/models"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type Migration struct {
	NewDB
	coa.DoSelf
}

func (ag *Migration) HandleError(ctx coa.Context, err error) error {
	panic(err.Error())
	return nil
}

func (ag *Migration) Do(ctx coa.Context) error {
	ag.DB().AutoMigrate(&ms.Msg{}, &ms.Pattern{})
	return nil
}
