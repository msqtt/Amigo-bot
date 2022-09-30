package bot

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mosqu1t0/Amigo-bot/utils/logcat"
)

type Bot struct {
	ws      *websocket.Conn
	HasBoot bool
	Info    QQInfo
}

var upgrader = websocket.Upgrader{}

func NewBot() *Bot {
	return &Bot{}
}

// 连接 websocket
func (bot *Bot) Start() {
	u := url.URL{
		Scheme: "ws",
		Host:   DefaultBotConfig.Ws.Addr,
		Path:   DefaultBotConfig.Ws.Path,
	}

	var senWs *websocket.Conn
	var err error
	for {
		logcat.Info("尝试连接到服务器: ", u.String())
		senWs, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			logcat.Error("连接服务器，失败: ", err)
			logcat.Info("bot 将于 3s 后重新连接...")
			time.Sleep(time.Second * 3)
		} else {
			break
		}
	}

	bot.ws = senWs
	bot.HasBoot = true
	logcat.Good("连接服务器，成功!")
	bot.showlogo()

	if DefaultBotConfig.Debug {
		logcat.Good("> Debug 模式开启, bot 将会打印原始 json <")
	} else {
		logcat.Good("> Debug 模式关闭 <")
	}

	//接收 lifecircle
	bot.receive()

	bot.getInfo()

	PluginMgr.finishInit()
}

// bot 开始接收消息
func (bot *Bot) Work() {
	if bot.HasBoot {
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

// 关闭websocket 连接
func (bot *Bot) Close() {
	err := bot.ws.WriteMessage(websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		logcat.ErrorEnd("无法发生关闭消息，即将强行关闭bot...")
	}
	bot.ws.Close()
	logcat.Good("bot 连接关闭成功!")
}

func (bot *Bot) showlogo() {
	splitLine := "----------------------------------------\n"
	version := ">           Amigo bot v1.1             <\n"
	fmt.Printf("\033[1;35m%s\033[0m\n",
		fmt.Sprint(botLogo, splitLine, version, splitLine),
	)
}

// 获取 bot 登录qq基本信息
func (bot *Bot) getInfo() {
	jsonBytes, err := bot.QuickTalk(GetBotInfo, &bot.Info)
	if err != nil {
		logcat.ErrorEnd("获取 bot 信息失败: ", err)
	}

	recv := struct {
		Data QQInfo `json:"data"`
	}{}
	err = json.Unmarshal(jsonBytes, &recv)
	if err != nil {
		logcat.ErrorEnd("bot 信息json 解析失败: ", err)
	}
	bot.Info = recv.Data
	logcat.Good(
		"bot 登录成功！登录账户 [QQ: ",
		bot.Info.UserId,
		" 昵称: ",
		bot.Info.NickName,
		"]",
	)
}
