package main

import (
	"github.com/mosqu1t0/Amigo-bot/bot"
	// 插件的引入要放在bo后
	_ "github.com/mosqu1t0/Amigo-bot/plugins/ezres"
)

func main() {
	b := bot.NewBot()
	b.Start()
	b.Work()
}
