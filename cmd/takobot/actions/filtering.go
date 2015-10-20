package actions

import (
	"fmt"
	"math/rand"

	eh "github.com/kazukgw/takobot/cmd/takobot/errorhandler"
	"github.com/kazukgw/takobot/cmd/takobot/msg"
	"github.com/kazukgw/takobot/cmd/takobot/store"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type Filtering struct {
	Func          func(*msg.Msg) bool
	WhiteChannel  []string
	BlackChannel  []string
	FromWhiteUser []string
	FromBlackUser []string
	ToWhiteUser   []string
	ToBlackUser   []string
	Ratio         float64
}

func (a *Filtering) Do(ctx coa.Context) error {
	if ag, ok := ctx.ActionGroup().(HasMsg); ok {
		m := ag.GetMsg()
		chanName := store.ChanByID(m.Channel).Name
		userName := store.UserByID(m.User).Name
		toUser := m.ToUser
		fmt.Printf(
			"Filter:%#v \nchanName: %#v\n userName: %#v\n toUser: %#v\n",
			a,
			chanName,
			userName,
			toUser,
		)
		if len(a.WhiteChannel) > 0 {
			for _, c := range a.WhiteChannel {
				if c == chanName {
					fmt.Println("filtering: white channel ok")
					return nil
				}
			}
			fmt.Println("filtering: white channel bad")
			return eh.ErrNonError
		}

		if len(a.BlackChannel) > 0 {
			for _, c := range a.BlackChannel {
				if c == chanName {
					fmt.Println("filtering: black channel bad")
					return eh.ErrNonError
				}
			}
			fmt.Println("filtering: black channel ok")
			return nil
		}

		if len(a.FromWhiteUser) > 0 {
			for _, u := range a.FromWhiteUser {
				if u == userName {
					fmt.Println("filtering: from white user ok")
					return nil
				}
			}
			fmt.Println("filtering: from white user bad")
			return eh.ErrNonError
		}

		if len(a.FromBlackUser) > 0 {
			for _, u := range a.FromBlackUser {
				if u == userName {
					fmt.Println("filtering: from black user bad")
					return eh.ErrNonError
				}
			}
			fmt.Println("filtering: from black user ok")
			return nil
		}

		if len(a.ToWhiteUser) > 0 {
			if toUser == nil {
				fmt.Println("filtering: to white user bad")
				return eh.ErrNonError
			}
			for _, u := range a.ToWhiteUser {
				if u == toUser.Name {
					fmt.Println("filtering: to white user ok")
					return nil
				}
			}
			fmt.Println("filtering: to white user bad")
			return eh.ErrNonError
		}

		if len(a.ToBlackUser) > 0 {
			if toUser == nil {
				fmt.Println("filtering: to white user bad")
				return eh.ErrNonError
			}
			for _, u := range a.ToBlackUser {
				if u == toUser.Name {
					fmt.Println("filtering: to black user bad")
					return eh.ErrNonError
				}
			}
			fmt.Println("filtering: to black user ok")
			return nil
		}

		if a.Ratio > 0.0 {
			if rand.Float64() <= a.Ratio {
				fmt.Println("filtering: ratio ok")
				return nil
			}
			fmt.Println("filtering: ratio bad")
			return eh.ErrNonError
		}

		if a.Func != nil && !a.Func(m) {
			fmt.Println("filtering: func bad")
			return eh.ErrNonError
		}
	}
	fmt.Println("filtering: ok")
	return nil
}
