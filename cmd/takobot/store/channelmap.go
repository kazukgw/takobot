package store

import (
	"github.com/kazukgw/takobot/Godeps/_workspace/src/github.com/nlopes/slack"
)

var IDChanMap = map[string]slack.Channel{}
var NameChanMap = map[string]slack.Channel{}

func MakeChanMap(chans []slack.Channel) {
	ResetChanMaps()
	for _, ch := range chans {
		IDChanMap[ch.ID] = ch
		NameChanMap[ch.Name] = ch
	}
}

func ResetChanMaps() {
	IDChanMap = map[string]slack.Channel{}
	NameChanMap = map[string]slack.Channel{}
}

func ChanByID(id string) slack.Channel {
	return IDChanMap[id]
}

func ChanByName(name string) slack.Channel {
	return NameChanMap[name]
}
