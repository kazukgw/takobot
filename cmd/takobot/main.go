package main

import "github.com/kazukgw/takobot/cron"

func main() {
	go ServeStats()
	go cron.KeepAlive("http://takobot.herokuapp.com")
	HandleEvent()
}
