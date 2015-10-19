package cron

import (
	"gopkg.in/robfig/cron.v2"
	"net/http"
)

func KeepAlive(addr string) {
	c := cron.New()
	c.AddFunc("@every 5m", func() { http.Get(addr) })
	c.Start()
	for {
		select {}
	}
}
