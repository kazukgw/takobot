package actiongroups

import (
	act "github.com/kazukgw/takobot/cmd/takobot/actions"
	"github.com/kazukgw/takobot/cmd/takobot/db"
	eh "github.com/kazukgw/takobot/cmd/takobot/errorhandler"
	"github.com/kazukgw/takobot/cmd/takobot/models"

	gq "github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/nlopes/slack"
)

type ReutersNews struct {
	act.Scrape
	coa.DoSelf
	act.SendMsg
}

func (ag ReutersNews) Schedule() string {
	return "0 15 03,07,11,23 * * *"
}

func (ag *ReutersNews) PreExec(ctx coa.Context) error {
	ag.Scrape.URL = "http://jp.reuters.com/news"
	return nil
}

func (ag *SendRegisteredMsg) Do(ctx coa.Context) error {
	text := ag.Document.Find(".column1 .columnLeft").Eq(0).Find("h2").Text()
	// $(".column1 .columnLeft").eq(0).find("h2 a").prop("href")

	ag.SendMsgWithPattern.Patterns = models.CompiledPatterns
	ag.SendMsgWithPattern.Source = ag.GetMsgFromCtx.Msg.Text
	return nil
}

func (ag *Nijihan) HandleError(ctx coa.Context, err error) error {
	fmt.Println(err.Error())
	return err
}
