package bot

import (
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
	"github.com/mosqu1t0/Amigo-bot/utils/logcat"
)

type Bot struct {
	ws          *websocket.Conn
	hasBoot     bool
	plugManager *PluginManager
}

var upgrader = websocket.Upgrader{}

func NewBot() *Bot {
	return &Bot{}
}

func (bot *Bot) Start() {
	u := url.URL{Scheme: "ws",
		Host: DefaultBotConfig.Ws.Addr,
		Path: DefaultBotConfig.Ws.Path}

	logcat.Info("尝试连接到服务器: ", u.String())
	senWs, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		logcat.ErrorEnd("连接服务器，失败: ", err)
	}

	bot.ws = senWs
	bot.hasBoot = true
	logcat.Good("连接服务器，成功!")
	InitPlugin(bot)
}

func (bot *Bot) Work() {
	if bot.hasBoot {
		logcat.Info("bot 开始工作啦")
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)
	Loop:
		for {
			select {
			case <-interrupt:
				logcat.Warn("bot 停止工作，关闭连接提醒...")
				bot.Close()
				break Loop
			default:
				bot.receive()
			}
		}
	}
	logcat.Warn("bot 没事可干, 水饺去了~")
}

func (bot *Bot) Close() {
	err := bot.ws.WriteMessage(websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		logcat.ErrorEnd("无法发生关闭消息，即将强行关闭bot...")
	}
	bot.ws.Close()
	logcat.Good("bot 连接关闭成功!")
}
