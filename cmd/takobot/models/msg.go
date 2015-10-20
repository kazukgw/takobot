package models

import (
	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/jinzhu/gorm"
)

type Msg struct {
	gorm.Model
	From     string `sql:"index;not null"`
	To       string `sql:"index"`
	Channel  string `sql:"index;not null"`
	FullBody string `sql:"size:1000"`
}
