package cron

import (
	"reflect"

	ags "github.com/kazukgw/takobot/cmd/takobot/actiongroups"
	ctxs "github.com/kazukgw/takobot/cmd/takobot/contexts"
	"github.com/kazukgw/takobot/cmd/takobot/log"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/nlopes/slack"
	"github.com/kazukgw/takobot/Godeps/_workspace/src/gopkg.in/robfig/cron.v2"
)

var crons = []interface{}{
	KeepAlive{},
	ags.Nijihan{},
	ags.Samishi{},
	ags.ReutersNews{},
}

type HasSchedule interface {
	Schedule() string
}

func Init(rtm *slack.RTM, client *slack.Client) {
	c := cron.New()
	for _, agSource := range crons {
		s := agSource.(HasSchedule).Schedule()
		if s != "" {
			_agSource := agSource
			c.AddFunc(s, func() {
				log.ActionGRP(reflect.TypeOf(_agSource).Name)
				ctxs.NewMsgContext(_agSource, nil, rtm, client).Exec()
			})
		}
	}
	c.Start()

	for {
		select {}
	}
}
