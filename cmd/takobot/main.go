package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/kazukgw/takobot/cron"
	"github.com/kazukgw/takobot/msghandler"
	"github.com/nlopes/slack"
)

func main() {
	go ServeStats()
	go cron.KeepAlive("http://takobot.herokuapp.com")
	go HandleEvent()
}
