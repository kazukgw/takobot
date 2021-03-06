package actiongroups

import (
	"math/rand"

	act "github.com/kazukgw/takobot/cmd/takobot/actions"
	"github.com/kazukgw/takobot/cmd/takobot/db"
	eh "github.com/kazukgw/takobot/cmd/takobot/errorhandler"
	"github.com/kazukgw/takobot/cmd/takobot/store"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/jinzhu/gorm"
	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type Samishi struct {
	db.NewDB
	act.MsgHistory
	coa.DoSelf
	act.GetRTMAndSendMsg

	eh.DefaultErrorHandler
}

func (ag Samishi) Schedule() string {
	return "0 45 02-13 * * *"
}

func (ag *Samishi) PreExec(ctx coa.Context) error {
	ag.MsgHistory.LastMinutes = 90
	ag.MsgHistory.Scope = func(db *gorm.DB) *gorm.DB {
		return db.Where("from_user != ?", store.UserByName("takobot").ID)
	}
	return nil
}

func (ag *Samishi) Do(ctx coa.Context) error {
	if len(ag.MsgHistory.Msgs) == 0 {
		msgs := []string{"暇や〜", "誰もおらんしさみしいわ"}
		ag.SendMsg.Msg = msgs[rand.Intn(len(msgs))]
		ag.SendMsg.Channel = "general"
	}
	return nil
}
