package actiongroups

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	act "github.com/kazukgw/takobot/cmd/takobot/actions"
	ctxs "github.com/kazukgw/takobot/cmd/takobot/contexts"
	"github.com/kazukgw/takobot/cmd/takobot/db"
	eh "github.com/kazukgw/takobot/cmd/takobot/errorhandler"
	"github.com/kazukgw/takobot/cmd/takobot/models"
	"github.com/kazukgw/takobot/cmd/takobot/store"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type AddPattern struct {
	db.NewDB
	act.GetMsgFromCtx
	act.Filtering
	coa.DoSelf
	act.GetRTMAndSendMsg

	eh.DefaultErrorHandler
}

var addPatternSplitArgs = regexp.MustCompile(`pattern[ ]+add:(.*)`)
var addPatternArgs = regexp.MustCompile(`(?:"([^"]*)")|(?:'([^']*)')`)

func getAddPatternArgs(source string) []string {
	ms := addPatternSplitArgs.FindAllStringSubmatch(source, -1)
	fmt.Printf("ms: %#v\n", ms)

	args := []string{}
	if len(ms) > 0 && len(ms[0]) > 1 {
		_ms := addPatternArgs.FindAllStringSubmatch(ms[0][1], -1)
		fmt.Printf("_ms: %#v\n", _ms)
		for _, m := range _ms {
			if len(m) > 2 {
				a := m[1]
				if a == "" {
					a = m[2]
				}
				if a != "" {
					args = append(args, a)
				}
			}
		}
	}
	return args
}

func (ag *AddPattern) PreExec(ctx coa.Context) error {
	ag.Filtering.ToPermitUser = []string{"takobot"}
	return nil
}

func (ag *AddPattern) Do(ctx coa.Context) error {
	mctx := ctx.(*ctxs.MsgContext)
	text := mctx.Msg.Text
	args := getAddPatternArgs(text)
	if len(args) < 1 {
		return errors.New(fmt.Sprintf("error: args = %v", args))
	}

	re, err := regexp.Compile(args[0])
	if err != nil {
		return err
	}

	ptn := &models.Pattern{Pattern: args[0], Strs: args[1:]}
	ag.DB().Create(ptn)
	if ag.DB().Error != nil {
		return ag.DB().Error
	}

	models.CompiledPatterns[re] = ptn.Strs

	ag.SendMsg.Msg = fmt.Sprintf("ﾊﾟｯﾀｰﾝ %s: %s", args[0], strings.Join(args[1:], ", "))
	ag.SendMsg.Channel = store.ChanByID(mctx.Msg.Channel).Name

	return nil
}

func (ag *AddPattern) HandleError(ctx coa.Context, err error) error {
	if err == eh.ErrNonError {
		return err
	}
	err = (&eh.DefaultErrorHandler{}).HandleError(ctx, err)
	mctx := ctx.(*ctxs.MsgContext)
	chanName := mctx.Msg.Channel
	rtm := mctx.RTM
	rtm.SendMessage(rtm.NewOutgoingMessage("Failed to add pattern.", store.ChanByName(chanName).ID))
	return err
}
