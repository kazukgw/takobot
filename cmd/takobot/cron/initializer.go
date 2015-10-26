package cron

import (
	"fmt"
	"reflect"

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
	c := cron.New()
	for _, agSource := range crons {
		s := agSource.(HasSchedule).Schedule()
		if s != "" {
			fmt.Printf("register action: %#v s:%#v", agSource, s)
			c.AddFunc(s, func() {
				fmt.Printf(
					"do the cron action group: %#v s: %#v",
					reflect.TypeOf(agSource).Name(),
					s,
				)
				ctxs.NewMsgContext(agSource, nil, rtm).Exec()
			})
		}
	}
	c.Start()

	for {
		select {}
	}
}
