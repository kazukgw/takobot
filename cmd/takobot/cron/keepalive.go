package cron

import (
	"net/http"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type KeepAlive struct {
	coa.DoSelf

	DefaultErrorHandler
}

func (ag KeepAlive) Schedule() string {
	return "@every 5m"
}

func (ag *KeepAlive) Do(ctx coa.Context) error {
	http.Get("http://takobot.herokuapp.com")
	return nil
}
