package actiongroups

import (
	"fmt"
	"math/rand"

	act "github.com/kazukgw/takobot/cmd/takobot/actions"
	"github.com/kazukgw/takobot/cmd/takobot/db"
	"github.com/kazukgw/takobot/cmd/takobot/store"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/jinzhu/gorm"
	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type Samishi struct {
	db.NewDB
	act.MsgHistory
	coa.DoSelf
	act.SendMsg
	db.CloseDB
}

func (ag Samishi) Schedule() string {
	return "0 40 02-13 * * *"
}

func (ag *Samishi) PreExec(ctx coa.Context) error {
	ag.MsgHistory.LastMinutes = 90
	ag.MsgHistory.Scope = func(db *gorm.DB) *gorm.DB {
		return db.Where("`from_user` != ?", store.UserByName("takobot").ID)
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

func (ag *Samishi) HandleError(ctx coa.Context, err error) error {
	fmt.Println(err.Error())
	return err
}
