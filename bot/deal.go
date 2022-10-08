package bot

import (
	"strings"

	"github.com/mosqu1t0/Amigo-bot/utils/logcat"
)

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

func (revm *RecvMessage) Analyze() ([]string, []map[string]string) {
	var resText []string
	var resMap []map[string]string
	dealWithCQ := func(cqStr string) (map[string]string, bool) {
		if len(cqStr) == 0 {
			return nil, false
		}
		cqSlice := strings.Split(cqStr, ",")
		if len(cqSlice) < 2 {
			return nil, false
		}
		resMap := make(map[string]string)
		resMap["CQ"] = strings.Split(cqSlice[0], ":")[1]

		for _, arg := range cqSlice[1:] {
			argSlice := strings.Split(arg, "=")
			resMap[argSlice[0]] = argSlice[1]
		}

		return resMap, true
	}

	msg := revm.Message

	lb := strings.Index(msg, "[")
	if lb != -1 && lb > 0 {
		resText = append(resText, msg[:lb])
	}
	for lb != -1 {
		rb := strings.Index(msg, "]")
		if cq, ok := dealWithCQ(msg[lb+1 : rb]); ok {
			resMap = append(resMap, cq)
		}
		msg = msg[rb+1:]
		lb = strings.Index(msg, "[")
		if lb != -1 && lb > 0 {
			resText = append(resText, msg[:lb])
		}
	}
	if len(msg) > 0 {
		resText = append(resText, msg)
	}

	return resText, resMap
}
