package actiongroups

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	act "github.com/kazukgw/takobot/cmd/takobot/actions"
	ctxs "github.com/kazukgw/takobot/cmd/takobot/contexts"
	"github.com/kazukgw/takobot/cmd/takobot/db"
	eh "github.com/kazukgw/takobot/cmd/takobot/errorhandler"
	"github.com/kazukgw/takobot/cmd/takobot/models"
	"github.com/kazukgw/takobot/cmd/takobot/store"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type ShowPattern struct {
	db.NewDB
	act.GetMsgFromCtx
	act.Filtering
	coa.DoSelf
	act.GetRTMAndSendMsg
	db.CloseDB
}

func (ag *ShowPattern) PreExec(ctx coa.Context) error {
	ag.Filtering.ToPermitUser = []string{"takobot"}
	return nil
}

var lsPatternSplitArgs = regexp.MustCompile(`pattern[ ]+ls:(.*)`)
var lsPatternArgs = regexp.MustCompile(`(\d+)`)

func getLsPatternArgs(source string) []string {
	ms := lsPatternSplitArgs.FindAllStringSubmatch(source, -1)
	fmt.Printf("ms: %#v\n", ms)

	args := []string{}
	if len(ms) > 0 && len(ms[0]) > 0 {
		_ms := lsPatternArgs.FindAllStringSubmatch(ms[0][1], -1)
		fmt.Printf("_ms: %#v\n", _ms)
		for _, m := range _ms {
			fmt.Printf("m: %#v\n", m[1])
			args = append(args, m[1])
		}
	}

	return args
}

func (ag *ShowPattern) Do(ctx coa.Context) error {
	mctx := ctx.(*ctxs.MsgContext)
	text := mctx.Msg.Text

	args := getLsPatternArgs(text)

	ps := []models.Pattern{}
	switch len(args) {
	case 0:
		ag.DB().Order("id asc").Find(&ps)
	case 1:
		l, _ := strconv.Atoi(args[0])
		ag.DB().Order("id asc").Limit(l).Find(&ps)
	case 2:
		l, _ := strconv.Atoi(args[0])
		p, _ := strconv.Atoi(args[1])
		ag.DB().Order("id asc").Limit(l).Offset(l * p).Find(&ps)
	}

	msgs := ""
	for _, p := range ps {
		msgs += fmt.Sprintf("#%v  /%v/  :  %v\n", p.ID, p.Pattern, strings.Join(p.Strs, ", "))
	}

	ag.SendMsg.Msg = msgs
	ag.SendMsg.Channel = store.ChanByID(mctx.Msg.Channel).Name

	return nil
}

func (ag *ShowPattern) HandleError(ctx coa.Context, err error) error {
	if err == eh.ErrNonError {
		return err
	}

	mctx := ctx.(*ctxs.MsgContext)
	fmt.Println(err.Error())
	chanName := mctx.Msg.Channel
	rtm := mctx.RTM
	rtm.SendMessage(rtm.NewOutgoingMessage("Failed to show pattern.", store.ChanByName(chanName).ID))
	return err
}
