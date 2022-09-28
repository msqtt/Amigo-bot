package ezres

import (
	"github.com/mosqu1t0/Amigo-bot/bot"
	"github.com/mosqu1t0/Amigo-bot/utils/logcat"
)

type Ezres struct {
	postType string
}

func init() {
	ezres := new(Ezres)
	ezres.postType = "request"
	bot.PluginMgr.AddPlugin(ezres) // 将插件手动添加到mgr

}

// 在PluginMgr之后的初始化函数，maybe放一些信息?
func (ezres *Ezres) Init() {
	logcat.Good("[远离害虫]，Ezres 插件加载完成！<3")
}

// 获取插件作用的消息类型
func (ezres *Ezres) GetType() string {
	return ezres.postType
}

// 实现插件interface, Action() 的是插件调用的主要功能
func (ezres *Ezres) Action(b *bot.Bot, v interface{}) {
	req, _ := v.(*bot.RecvRequest) // 断言转换为指针类型
	// 具体逻辑
	ifAgree := false
	for _, root := range bot.DefaultBotConfig.Root {
		if root == req.UserId {
			ifAgree = true
			break
		}
	}
	switch req.RequestType {
	case bot.GruRequestType:
		b.Send(bot.GruRequestApi, struct {
			Flag    string `json:"flag"`
			SubType string `json:"sub_type"`
			Approve bool   `json:"approve"`
			Reason  string `json:"reason"`
		}{req.Flag, req.SubType, ifAgree, "我不跟陌生人走嗷"})
	case bot.FriRequestType:
		b.Send(bot.FriRequestApi, struct {
			Flag    string `json:"flag"`
			Approve bool   `json:"approve"`
		}{req.Flag, ifAgree})
	default:
		logcat.Error("哈？未知的邀请格式...")
		return
	}
	if ifAgree {
		logcat.Good("bot 已接受 [QQ: ", req.UserId, "] 的邀请: ", req.GroupId)
	} else {
		logcat.Good("bot 拒绝 [QQ: ", req.UserId, "] 的邀请: ", req.GroupId)
	}
}
