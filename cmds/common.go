package cmds

import (
	"math/rand"

	"strconv"
)

func fortunate() string {
	randNumber := rand.Intn(100)
	switch {
	case randNumber > 90:
		return "你再抽一次🤡"
	case randNumber > 85:
		return "仙草吉👑"
	case randNumber > 75:
		return "大吉⭐"
	case randNumber > 60:
		return "中吉🌟"
	case randNumber > 40:
		return "小吉👌"
	case randNumber > 20:
		return "末吉🙏"
	case randNumber > 5:
		return "凶💀"
	default:
		return "大凶☠️"
	}
}

func draw() string {
	return strconv.Itoa(rand.Intn(11) + 1)
}
