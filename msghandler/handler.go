package msghandler

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func AddHandler(pattern string) bool {
	return true
}

func Handle(pattern string) string {
	fmt.Println("pattern: " + pattern)
	msgs := strings.Split(pattern, ":")
	if len(msgs) > 1 {
		return patternMatch(msgs[1])
	}
	return ""
}

var mapPattern = map[string][]string{
	"名前は":    {"takobotやで", "しっとるやろ"},
	"つらい":    {"なにがあったんや", "はなし、きくで"},
	"たのしい":   {"わいも楽しいわ", "ええなぁ。うらやましいわ。"},
	"なにしてん":  {"まぁいろいろしてるわ", "俺も意外と忙しいねんで"},
	"なにしてんの": {"まぁいろいろしてるわ", "今日は暇やったわ"},
	"おい":     {"なんや", "うるさい", "おいとは何ごとや"},
	"なぁなぁ":   {"なんや", "だまれ", "うん？"},
	"なぁ":     {"もうええで", "うんうん"},
	"どない":    {"ええかんじやわ", "いや〜、ええわ〜"},
	"あつい":    {"いや〜、あついで", "俺、機械やからなぁ、、、", "気温とかわからんわ"},
	"さむい":    {"ほんま？俺あついで", "俺、機械やからわからんわ"},
	"あそぼ":    {"すまんな。いま忙しいわ。", "俺、何して遊んだええねん"},
	"おっぱい":   {"おっぱいおっぱい", "やわらかいよな〜"},
	"すき":     {"おっぱいがすきなんやろ", "ちょっと重いわ"},
	"まんまん":   {"あほか", "シモネタとかちょっと引きますわ"},
}

func patternMatch(pattern string) string {
	var msgs []string
	for k, ms := range mapPattern {
		if strings.Contains(pattern, k) {
			msgs = ms
			break
		}
	}

	if len(msgs) == 0 {
		msgs = []string{"わからんわ"}
	}

	rand.Seed(time.Now().UnixNano())
	return msgs[rand.Intn(len(msgs))]
}
