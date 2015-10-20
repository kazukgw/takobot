package main

import (
	"fmt"

	ctxs "github.com/kazukgw/takobot/cmd/takobot/contexts"
	"github.com/kazukgw/takobot/cmd/takobot/db"
	"github.com/kazukgw/takobot/cron"
)

func main() {
	fmt.Println("hogehoge")
	ctxs.NewContext(&db.Migration{}).Exec()
	go ServeStats()
	go cron.KeepAlive("http://takobot.herokuapp.com")
	HandleEvent()
}
