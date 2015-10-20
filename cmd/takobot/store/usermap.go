package store

import (
	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/nlopes/slack"
)

var IDUserMap = map[string]slack.User{}
var NameUserMap = map[string]slack.User{}

func MakeUserMap(users []slack.User) {
	for _, u := range users {
		IDUserMap[u.ID] = u
		NameUserMap[u.Name] = u
	}
}

func UserByID(id string) slack.User {
	return IDUserMap[id]
}

func UserByName(name string) slack.User {
	return NameUserMap[name]
}
