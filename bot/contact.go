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

	var reckind RecvKind
	json.Unmarshal(bytes, &reckind)

	switch reckind.PostType {
	case MsgPostType:
		msg := new(RecvMessage)
		json.Unmarshal(bytes, msg)
		bot.handleMessage(msg)
	case ReqPostType:
		req := new(RecvRequest)
		json.Unmarshal(bytes, req)
		bot.handleRequest(req)
	case NtsPostType:
		nts := new(RecvNotice)
		json.Unmarshal(bytes, nts)
		bot.handleNotise(nts)
	case MtaPostType:
		mta := new(RecvMeta)
		json.Unmarshal(bytes, mta)
		bot.handleMeta(mta)
	}
}

func (bot *Bot) Send(action string, v interface{}) {
	err := bot.ws.WriteJSON(
		SendRespondJson{Action: action, Params: v, Echo: echo})
	if err != nil {
		logcat.Error("发送Json失败: ", err)
	}
}
