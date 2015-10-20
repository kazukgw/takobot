package actiongroups

import (
	"fmt"
	"regexp"

	act "github.com/kazukgw/takobot/cmd/takobot/actions"
	ctxs "github.com/kazukgw/takobot/cmd/takobot/contexts"
	"github.com/kazukgw/takobot/cmd/takobot/db"
	eh "github.com/kazukgw/takobot/cmd/takobot/errorhandler"
	"github.com/kazukgw/takobot/cmd/takobot/models"
	"github.com/kazukgw/takobot/cmd/takobot/store"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type RemovePattern struct {
	db.NewDB
	act.GetMsgFromCtx
	act.Filtering
	coa.DoSelf
	act.SendMsg
	db.CloseDB
}

func (ag *RemovePattern) PreExec(ctx coa.Context) error {
	ag.Filtering.ToWhiteUser = []string{"takobot"}
	return nil
}

var rmPatternSplitArgs = regexp.MustCompile(`pattern[ ]+rm:(.*)`)
var rmPatternArgs = regexp.MustCompile(`(\d+)`)

func getRmPatternArgs(source string) []string {
	ms := rmPatternSplitArgs.FindAllStringSubmatch(source, -1)
	fmt.Printf("ms: %#v\n", ms)

	args := []string{}
	if len(ms) > 0 && len(ms[0]) > 0 {
		_ms := rmPatternArgs.FindAllStringSubmatch(ms[0][1], -1)
		fmt.Printf("_ms: %#v\n", _ms)
		for _, m := range _ms {
			fmt.Printf("m: %#v\n", m[1])
			args = append(args, m[1])
		}
	}

	return args
}

func (ag *RemovePattern) Do(ctx coa.Context) error {
	mctx := ctx.(*ctxs.MsgContext)
	text := mctx.Msg.Text

	args := getRmPatternArgs(text)

	ag.DB().Where("id in (?)", args).Delete(models.Pattern{})

	ag.SendMsg.Msg = fmt.Sprint("ﾊﾟｯﾀｰﾝ ﾘﾑｰﾌﾞ!!")
	ag.SendMsg.Channel = store.ChanByID(mctx.Msg.Channel).Name

	return nil
}

func (ag *RemovePattern) HandleError(ctx coa.Context, err error) error {
	if err == eh.ErrNonError {
		return err
	}

	mctx := ctx.(*ctxs.MsgContext)
	fmt.Println(err.Error())
	chanName := mctx.Msg.Channel
	rtm := mctx.RTM
	rtm.SendMessage(rtm.NewOutgoingMessage("Failed to remove pattern.", store.ChanByName(chanName).ID))
	return err
}
