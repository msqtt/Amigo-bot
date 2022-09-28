package bot

import "github.com/mosqu1t0/Amigo-bot/utils/logcat"

func (bot *Bot) handleMessage(msg *RecvMessage) {
	logcat.Info("收到了消息: ", msg.String())
	for _, p := range PluginMgr.plugins[MsgPostType] {
		go p.Action(bot, msg)
	}
}

func (bot *Bot) handleRequest(req *RecvRequest) {
	logcat.Info("收到了请求: ", req.String())
	for _, p := range PluginMgr.plugins[ReqPostType] {
		go p.Action(bot, req)
	}
}

func (bot *Bot) handleNotise(nts *RecvNotice) {
	logcat.Info("收到了信息: ", nts.String())
	for _, p := range PluginMgr.plugins[NtsPostType] {
		go p.Action(bot, nts)
	}
}

func (bot *Bot) handleMeta(mta *RecvMeta) {
	if mta.MetaEvenType != hertMetaType {
		logcat.Info("收到了元信息: ", mta.String())
	}
	for _, p := range PluginMgr.plugins[MtaPostType] {
		go p.Action(bot, mta)
	}
}
