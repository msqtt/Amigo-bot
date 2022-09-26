package bot

import (
	"encoding/json"

	"github.com/mosqu1t0/Amigo-bot/utils/logcat"
)

func (bot *Bot) receive() {
	_, bytes, err := bot.ws.ReadMessage()
	if err != nil {
		logcat.Error("Read bytes from ws error: ", err)
	}

	var reckind RecKind
	json.Unmarshal(bytes, &reckind)

	switch reckind.PostType {
	case MsgPostType:
		msg := new(RecMessage)
		json.Unmarshal(bytes, msg)
		bot.handleMessage(msg)
	case ReqPostType:
		req := new(RecRequest)
		json.Unmarshal(bytes, req)
		bot.handleRequest(req)
	case NtsPostType:
		nts := new(RecNotice)
		json.Unmarshal(bytes, nts)
		bot.handleNotise(nts)
	case MtaPostType:
		mta := new(RecMeta)
		json.Unmarshal(bytes, mta)
		bot.handleMeta(mta)
	}
}

func (bot *Bot) send(action string, v interface{}) {
	str, _ := json.Marshal(SendRespondJson{
		Action: action,
		Params: v,
		Echo:   echo,
	})
	logcat.Info("发送消息...", string(str))
	err := bot.ws.WriteJSON(
		SendRespondJson{Action: action, Params: v, Echo: echo})
	if err != nil {
		logcat.Error("发送Json失败: ", err)
	}
}
