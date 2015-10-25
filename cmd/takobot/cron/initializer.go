package cron

import (
	ags "github.com/kazukgw/takobot/cmd/takobot/actiongroups"
	ctxs "github.com/kazukgw/takobot/cmd/takobot/contexts"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/nlopes/slack"
	"github.com/kazukgw/takobot/Godeps/_workspace/src/gopkg.in/robfig/cron.v2"
)

var crons = []interface{}{
	KeepAlive{},
	ags.Nijihan{},
	ags.Samishi{},
}

type HasSchedule interface {
	Schedule() string
}

func Init(rtm *slack.RTM) {
	for _, action := range crons {
		c := cron.New()
		s := action.(HasSchedule).Schedule()
		if s != "" {
			c.AddFunc(s, func() {
				ctxs.NewMsgContext(action, nil, rtm).Exec()
			})
		}
	}

	for {
		select {}
	}
}
