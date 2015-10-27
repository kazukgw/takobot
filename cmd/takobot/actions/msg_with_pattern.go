package actions

import (
	"math/rand"
	"regexp"
	"time"

	"github.com/kazukgw/takobot/cmd/takobot/log"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type MsgWithPattern struct {
	Patterns  map[*regexp.Regexp][]string
	Source    string
	ResultMsg string
}

func (a *MsgWithPattern) Do(ctx coa.Context) error {
	log.Action("msg with pattern ==>")
	msgs := make([]string, 0)
	for re, ms := range a.Patterns {
		if re.Match([]byte(a.Source)) {
			msgs = append(msgs, ms...)
		}
	}
	if len(msgs) == 0 {
		msgs = []string{"すまん。わからんわ", "ちょいむずかしいわ"}
	}
	rand.Seed(time.Now().UnixNano())
	a.ResultMsg = msgs[rand.Intn(len(msgs))]
	return nil
}
