package bot

import "github.com/mosqu1t0/Amigo-bot/utils/logcat"

func (bot *Bot) handleMessage(msg *RecMessage) {
	logcat.Info("收到了消息: ", msg.String())
}

func (bot *Bot) handleRequest(req *RecRequest) {
	logcat.Info("收到了请求: ", req.String())

	ifAgree := false
	for _, root := range DefaultBotConfig.Root {
		logcat.Print(root)
		if root == req.UserId {
			ifAgree = true
			break
		}
	}
	switch req.RequestType {
	case GruRequestType:
		bot.send(GruRequestApi, struct {
			Flag    string `json:"flag"`
			SubType string `json:"sub_type"`
			Approve bool   `json:"approve"`
			Reason  string `json:"reason"`
		}{req.Flag, req.SubType, ifAgree, "我不跟陌生人走嗷"})
	case FriRequestType:
		bot.send(FriRequestApi, struct {
			Flag    string `json:"flag"`
			Approve bool   `json:"approve"`
		}{req.Flag, ifAgree})
	default:
		logcat.Error("哈？未知的邀请格式...")
	}
}

func (bot *Bot) handleNotise(nts *RecNotice) {
	logcat.Info("收到了信息: ", nts.String())
}

func (bot *Bot) handleMeta(mta *RecMeta) {
	if mta.MetaEvenType != hertMetaType {
		logcat.Info("收到了元信息: ", mta.String())
	}
}
