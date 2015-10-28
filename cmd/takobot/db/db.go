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

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open("postgres", DBURL())
	if err != nil {
		panic(err.Error())
	}
	DB = &db
}

type NewDB struct {
	db *gorm.DB
}

type HasDB interface {
	DB() *gorm.DB
}

func (a *NewDB) Do(ctx coa.Context) error {
	a.db = DB
	return nil
}

func (a *NewDB) DB() *gorm.DB {
	return a.db
}
