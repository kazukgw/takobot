package cron

import (
	"fmt"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
)

type DefaultErrorHandler struct {
}

func (eh *DefaultErrorHandler) HandleError(ctx coa.Context, err error) error {
	fmt.Println(err.Error())
	return err
}
