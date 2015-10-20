package main

import (
	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa/web"
	"github.com/zenazn/goji"
)

func main() {
	hb := web.GojiHandlerBuilder{NewContext}

	goji.Get("/user", hb.Build(GetUser{}))
	goji.Post("/user", hb.Build(UpdateUser{}))
	goji.Serve()
}
