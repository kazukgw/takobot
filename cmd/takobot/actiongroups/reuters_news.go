package actiongroups

import (
	"fmt"

	act "github.com/kazukgw/takobot/cmd/takobot/actions"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/nlopes/slack"
)

type ReutersNews struct {
	act.Scrape
	coa.DoSelf
	act.GetClient
	act.SendAttachments
}

func (ag ReutersNews) Schedule() string {
	return "0 15 03,07,11,23 * * *"
	// return "@every 1m"
}

func (ag *ReutersNews) PreExec(ctx coa.Context) error {
	ag.Scrape.URL = "http://jp.reuters.com/news"
	return nil
}

func (ag *ReutersNews) Do(ctx coa.Context) error {
	feature := ag.Document.Find(".column1 .columnLeft").Eq(0).Find(".feature")
	title := feature.Find("h2").Text()
	text := feature.Find("p").Text()
	href, _ := feature.Find("h2 a").Attr("href")
	fmt.Printf("title: %v\ntext: %v\nhref: %v\n", title, text, href)
	ag.SendAttachments.Attachment = slack.Attachment{
		Pretext:   "ニュース持ってきたやで〜",
		Title:     title,
		TitleLink: "http://jp.reuters.com/news" + href,
		Text:      text,
	}
	return nil
}

func (ag *ReutersNews) HandleError(ctx coa.Context, err error) error {
	fmt.Println(err.Error())
	return err
}
