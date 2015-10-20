package db

import (
	"os"

	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/jinzhu/gorm"
	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/kazukgw/coa"
	_ "github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/lib/pq"
)

func DBURL() string {
	return os.Getenv("DATABASE_URL")
}

type NewDB struct {
	db *gorm.DB
}

type HasDB interface {
	DB() *gorm.DB
}

func (a *NewDB) Do(ctx coa.Context) error {
	db, err := gorm.Open("postgres", DBURL())
	if err != nil {
		return err
	}
	a.db = &db
	return nil
}

func (a *NewDB) DB() *gorm.DB {
	return a.db
}

type CloseDB struct {
}

func (a *CloseDB) Do(ctx coa.Context) error {
	ag := ctx.ActionGroup()
	ag.(HasDB).DB().Close()
	return nil
}
