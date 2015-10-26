package actions

import (
	"fmt"

	ctxs "github.com/kazukgw/takobot/cmd/takobot/contexts"
	"github.com/kazukgw/takobot/cmd/takobot/store"

	gq "github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type Scrape struct {
	URL string
	*gq.Document
}

func (a *Scrape) Do(ctx coa.Context) error {
	a.Document = gq.NewDocument(a.URL)
	return nil
}
