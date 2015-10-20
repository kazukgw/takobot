package models

import (
	"database/sql/driver"
	"encoding/json"
	"regexp"

	"github.com/jinzhu/gorm"
)

type Pattern struct {
	gorm.Model
	Pattern string `sql:"not null"`
	Strs    `sql:"type:text;not null"`
}

type Patterns []*Pattern

var CompiledPatterns = map[*regexp.Regexp][]string{}

type Strs []string

func (l Strs) Value() (driver.Value, error) {
	return json.Marshal(l)
}

func (l *Strs) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), l)
}
