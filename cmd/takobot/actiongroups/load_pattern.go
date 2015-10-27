package actiongroups

import (
	"regexp"

	"github.com/kazukgw/takobot/cmd/takobot/db"
	eh "github.com/kazukgw/takobot/cmd/takobot/errorhandler"
	"github.com/kazukgw/takobot/cmd/takobot/models"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type LoadPattern struct {
	db.NewDB
	coa.DoSelf

	models.Patterns
	eh.DefaultErrorHandler
}

func (ag *LoadPattern) Do(ctx coa.Context) error {
	ag.DB().Find(&ag.Patterns)
	for _, p := range ag.Patterns {
		re := regexp.MustCompile(p.Pattern)
		models.CompiledPatterns[re] = p.Strs
	}
	return nil
}
