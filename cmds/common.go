package cmds

import (
	"math/rand"

	"strconv"
)

var (
	fortunateArr = [8]string{"大凶", "凶", "末吉", "小吉", "中吉", "大吉", "仙草吉", "你再抽一次( ºωº )"}
)

func fortunate() string {
	return fortunateArr[rand.Intn(6)]
}

func draw() string {
	return strconv.Itoa(rand.Intn(11) + 1)
}
