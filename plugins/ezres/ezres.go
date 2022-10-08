/*
一个非常简单的请求处理插件，拒绝除了root qq 以外的好友或群请求
*/
package ezres

import (
	"github.com/mosqu1t0/Amigo-bot/bot"
	"github.com/mosqu1t0/Amigo-bot/utils/logcat"
)

/*
插件必须实现的接口

	type Plugin interface {
	    Init()
	    GetType() string
	    Action(b *bot.Bot, v interface{})
	}
*/
type Ezres struct {
}

// 必须在 init 将插件添加到 PluginMgr 中，否则插件无法被读取
// 为了执行init, 必须在 main.go 中import 该插件包
func init() {
	ezres := new(Ezres)
	bot.PluginMgr.AddPlugin(ezres)
}

// 在 PluginMgr 之后的初始化函数，maybe 放一些信息?
func (ezres *Ezres) Init() {
	logcat.Good("[远离害虫]，Ezres 插件加载完成！<3")
}

// 获取插件作用的消息类型
func (ezres *Ezres) GetType() string {
	return bot.ReqPostType
}

// 实现插件的interface, Action() 放插件的主要功能
func (ezres *Ezres) Action(b *bot.Bot, v interface{}) {
	req, _ := v.(*bot.RecvRequest) // 记得断言为指针类型
	isAgree := false
	for _, root := range bot.DefaultBotConfig.Root {
		if root == req.UserId {
			isAgree = true
			break
		}
	}
	switch req.RequestType {

	case bot.GruRequestType:
		b.Send(
			bot.GruRequestApi,
			struct {
				Flag    string `json:"flag"`
				SubType string `json:"sub_type"`
				Approve bool   `json:"approve"`
				Reason  string `json:"reason"`
			}{req.Flag, req.SubType, isAgree, "我不跟陌生人走嗷"})

	case bot.FriRequestType:
		b.Send(
			bot.FriRequestApi,
			struct {
				Flag    string `json:"flag"`
				Approve bool   `json:"approve"`
			}{req.Flag, isAgree})

	default:
		logcat.Error("哈？未知的邀请格式...")
		return
	}
	if isAgree {
		logcat.Good("bot 已接受 [QQ: ", req.UserId, "] 的邀请: ", req.GroupId)
	} else {
		logcat.Good("bot 拒绝 [QQ: ", req.UserId, "] 的邀请: ", req.GroupId)
	}
}

// 关闭 bot 前时执行的，销毁插件的方法
func (ezres Ezres) Destroy() {
}
