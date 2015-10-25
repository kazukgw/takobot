package main

import (
	"fmt"

	ctxs "github.com/kazukgw/takobot/cmd/takobot/contexts"
	"github.com/kazukgw/takobot/cmd/takobot/db"
)

func main() {
	fmt.Println("hogehoge")
	ctxs.NewContext(&db.Migration{}).Exec()
	go ServeStats()
	HandleEvent()
}
