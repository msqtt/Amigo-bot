package main

import (
	"github.com/mosqu1t0/Amigo-bot/bot"
)

func main() {
	b := bot.NewBot()
	b.Start()
	b.Work()
}
