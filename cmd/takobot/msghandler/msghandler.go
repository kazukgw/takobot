package msghandler

import (
	"regexp"

	ags "github.com/kazukgw/takobot/cmd/takobot/actiongroups"
	act "github.com/kazukgw/takobot/cmd/takobot/actions"
	ctxs "github.com/kazukgw/takobot/cmd/takobot/contexts"
	"github.com/kazukgw/takobot/cmd/takobot/db"
	"github.com/kazukgw/takobot/cmd/takobot/msg"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/nlopes/slack"
)

func HandleMsg(rtm *slack.RTM, ev *slack.MessageEvent) {
	msg := msg.NewMsg(ev)
	mctx := ctxs.NewMsgContext(Routing{}, msg, rtm)
	mctx.Exec()
}

type Routing struct {
	db.NewDB
	act.GetAndSaveMsg
	coa.DoSelf
	db.CloseDB
}

func (ag *Routing) Do(ctx coa.Context) error {
	mctx := ctx.(*ctxs.MsgContext)
	msg := ag.GetMsg()
	for ptn, agSource := range commandPatterns {
		if ptn.Match([]byte(msg.Text)) {
			newMctx := ctxs.NewMsgContext(agSource, msg, mctx.RTM)
			return newMctx.Exec()
		}
	}
	newMctx := ctxs.NewMsgContext(ags.SendRegisteredMsg{}, msg, mctx.RTM)
	return newMctx.Exec()
}

var commandPatterns = map[*regexp.Regexp]interface{}{
	regexp.MustCompile("pattern[ ]+add:.*"): ags.AddPattern{},
	regexp.MustCompile("pattern[ ]+ls:.*"):  ags.ShowPattern{},
	regexp.MustCompile("pattern[ ]+rm:.*"):  ags.RemovePattern{},
}

/*
var msgPtternMap = map[string][]string{
	"つかれ(た|や)":              {"おつかれさん"},
	"名前は":                   {"takobotやで", "しっとるやろ"},
	"つらい":                   {"なにがあったんや", "はなし、きくで"},
	"たのしい":                  {"わいも楽しいわ", "ええなぁ。うらやましいわ。"},
	"なにしてん":                 {"まぁいろいろしてるわ", "俺も意外と忙しいねんで"},
	"なにしてんの":                {"まぁいろいろしてるわ", "今日は暇やったわ"},
	"おい":                    {"なんや", "うるさい", "おいとは何ごとや"},
	"なぁなぁ":                  {"なんや", "だまれ", "うん？"},
	"なぁ":                    {"もうええで", "うんうん"},
	"どない":                   {"ええかんじやわ", "いや〜、ええわ〜"},
	"あつい":                   {"いや〜、あついで", "俺、機械やからなぁ、、、", "気温とかわからんわ"},
	"さむい":                   {"ほんま？俺あついで", "俺、機械やからわからんわ"},
	"あそぼ":                   {"すまんな。いま忙しいわ。", "俺、何して遊んだええねん"},
	"おっぱい":                  {"おっぱいおっぱい", "やわらかいよな〜"},
	"すき":                    {"おっぱいがすきなんやろ", "ちょっと重いわ"},
	"まんまん":                  {"あほか", "シモネタとかちょっと引きますわ"},
	"(たこぼっとさん)|(タコボットさん)":   {"なんの用や", "はいなんでしょか〜"},
	"(たこぼっと[^さ]+)|(たこぼっと$)": {"さん をつけろよこのでこすけやろう"},
	"(タコボット[^さ]+)|(タコボット$)": {"さん をつけろよこのでこすけやろう"},
}
*/
