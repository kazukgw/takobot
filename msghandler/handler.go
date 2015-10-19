package msghandler

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func AddHandler(pattern string) bool {
	return true
}

func Handle(msg string) string {
	msgs := strings.Split(msg, ":")
	if len(msgs) > 1 {
		return patternMatch(msgs[1])
	}
	return ""
}

var mapPattern = map[string][]string{
	"つかれ(た|や)":              {"おつかれさん"},
	"名前は":                   {"takobotやで", "しっとるやろ"},
	"つらい":                   {"なにがあったんや", "はなし、きくで"},
	"たのしい":                  {"わいも楽しいわ", "ええなぁ。うらやましいわ。"},
	"なにしてん":                 {"まぁいろいろしてるわ", "俺も意外と忙しいねんで"},
	"なにしてんの":                {"まぁいろいろしてるわ", "今日は暇やったわ"},
	"おい":                    {"なんや", "うるさい", "おいとは何ごとや"},
	"なぁなぁ":                  {"なんや", "だまれ", "うん？"},
	"なぁ":                    {"もうええで", "うんうん"},
	"どない":                   {"ええかんじやわ", "いや〜、ええわ〜"},
	"あつい":                   {"いや〜、あついで", "俺、機械やからなぁ、、、", "気温とかわからんわ"},
	"さむい":                   {"ほんま？俺あついで", "俺、機械やからわからんわ"},
	"あそぼ":                   {"すまんな。いま忙しいわ。", "俺、何して遊んだええねん"},
	"おっぱい":                  {"おっぱいおっぱい", "やわらかいよな〜"},
	"すき":                    {"おっぱいがすきなんやろ", "ちょっと重いわ"},
	"まんまん":                  {"あほか", "シモネタとかちょっと引きますわ"},
	"(たこぼっとさん)|(タコボットさん)":   {"なんの用や", "はいなんでしょか〜"},
	"(たこぼっと[^さ]+)|(たこぼっと$)": {"さん をつけろよこのでこすけやろう"},
	"(タコボット[^さ]+)|(タコボット$)": {"さん をつけろよこのでこすけやろう"},
}

func patternMatch(msg string) string {
	var msgs []string
	for k, ms := range mapPattern {
		regex := regexp.MustCompile(k)
		if regex.Match([]byte(msg)) {
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
