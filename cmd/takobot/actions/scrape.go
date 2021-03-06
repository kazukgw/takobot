package actions

import (
	"github.com/kazukgw/takobot/cmd/takobot/log"

	gq "github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type Scrape struct {
	URL string
	*gq.Document
}

func (a *Scrape) Do(ctx coa.Context) error {
	log.Action("==> scrape")
	doc, err := gq.NewDocument(a.URL)
	if err != nil {
		return err
	}
	a.Document = doc
	return nil
}
