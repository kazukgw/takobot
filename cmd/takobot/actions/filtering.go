package actions

import (
	"fmt"
	"math/rand"

	eh "github.com/kazukgw/takobot/cmd/takobot/errorhandler"
	"github.com/kazukgw/takobot/cmd/takobot/log"
	"github.com/kazukgw/takobot/cmd/takobot/msg"
	"github.com/kazukgw/takobot/cmd/takobot/store"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type Filtering struct {
	Func           func(*msg.Msg) bool
	PermitChannel  []string
	DenyChannel    []string
	FromPermitUser []string
	FromDenyUser   []string
	ToPermitUser   []string
	ToDenyUser     []string
	Ratio          float64
}

func (a *Filtering) Do(ctx coa.Context) error {
	log.Action("filtering ==>")
	if ag, ok := ctx.ActionGroup().(HasMsg); ok {
		m := ag.GetMsg()
		chanName := store.ChanByID(m.Channel).Name
		userName := store.UserByID(m.User).Name
		toUser := m.ToUser
		if len(a.PermitChannel) > 0 {
			for _, c := range a.PermitChannel {
				if c == chanName {
					log.Info("filtering: permit channel ok")
					return nil
				}
			}
			log.Info("filtering: permit channel bad")
			return eh.ErrNonError
		}

		if len(a.DenyChannel) > 0 {
			for _, c := range a.DenyChannel {
				if c == chanName {
					log.Info("filtering: deny channel bad")
					return eh.ErrNonError
				}
			}
			log.Info("filtering: deny channel ok")
			return nil
		}

		if len(a.FromPermitUser) > 0 {
			for _, u := range a.FromPermitUser {
				if u == userName {
					log.Info("filtering: from permit user ok")
					return nil
				}
			}
			log.Info("filtering: from permit user bad")
			return eh.ErrNonError
		}

		if len(a.FromDenyUser) > 0 {
			for _, u := range a.FromDenyUser {
				if u == userName {
					log.Info("filtering: from deny user bad")
					return eh.ErrNonError
				}
			}
			log.Info("filtering: from deny user ok")
			return nil
		}

		if len(a.ToPermitUser) > 0 {
			if toUser == nil {
				log.Info("filtering: to permit user bad")
				return eh.ErrNonError
			}
			for _, u := range a.ToPermitUser {
				if u == toUser.Name {
					log.Info("filtering: to permit user ok")
					return nil
				}
			}
			log.Info("filtering: to permit user bad")
			return eh.ErrNonError
		}

		if len(a.ToDenyUser) > 0 {
			if toUser == nil {
				log.Info("filtering: to permit user bad")
				return eh.ErrNonError
			}
			for _, u := range a.ToDenyUser {
				if u == toUser.Name {
					log.Info("filtering: to deny user bad")
					return eh.ErrNonError
				}
			}
			log.Info("filtering: to deny user ok")
			return nil
		}

		if a.Ratio > 0.0 {
			if rand.Float64() <= a.Ratio {
				log.Info("filtering: ratio ok")
				return nil
			}
			log.Info("filtering: ratio bad")
			return eh.ErrNonError
		}

		if a.Func != nil && !a.Func(m) {
			log.Info("filtering: func bad")
			return eh.ErrNonError
		}
	}
	log.Info("filtering: ok")
	return nil
}
