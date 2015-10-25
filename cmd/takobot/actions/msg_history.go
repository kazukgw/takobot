package actions

import (
	"time"

	"github.com/kazukgw/takobot/cmd/takobot/db"
	"github.com/kazukgw/takobot/cmd/takobot/models"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/jinzhu/gorm"
	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type MsgHistory struct {
	LastMinutes int
	LastMsgNum  int
	Scope       func(*gorm.DB) *gorm.DB
	Msgs        []*models.Msg
}

func (a *MsgHistory) Do(ctx coa.Context) error {
	ag := ctx.ActionGroup()
	db := ag.(db.HasDB).DB()
	if a.Scope != nil {
		db = db.Scopes(a.Scope)
	}
	if a.LastMinutes > 0 {
		db.Where(
			"created_at > ?",
			time.Now().Add(time.Duration(-1*a.LastMinutes)*time.Minute),
		).Find(&a.Msgs)
		return nil
	}

	if a.LastMsgNum > 0 {
		db.Order("id desc").Limit(a.LastMsgNum).Find(&a.Msgs)
	}
	return nil
}
