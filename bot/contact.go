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

	if DefaultBotConfig.Debug {
		logcat.Debug(string(bytes))
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

// 发送请求, 不处理响应信息, 如果发送失败会返回error
// 第一个参数是api的url，第二个参数是请求的 params struct
func (bot *Bot) Send(action string, v interface{}) error {

	bytes, _ := json.Marshal(SendRespondJson{Action: action, Params: v, Echo: echo})

	if DefaultBotConfig.Debug {
		logcat.Debug(string(bytes))
	}
	err := bot.ws.WriteJSON(
		SendRespondJson{Action: action, Params: v, Echo: echo},
	)
	if err != nil {
		return err
	}
	return nil
}

// 快速反应，发送请求，接收json 的 []byte, 如果有错误会返回error
// 第一个参数是api的url，第二个参数是请求的 params struct
func (bot *Bot) QuickTalk(action string, v interface{}) ([]byte, error) {
	errSend := bot.Send(action, v)
	if errSend != nil {
		return nil, errSend
	}
	_, bytes, errRead := bot.ws.ReadMessage()
	if errRead != nil {
		return nil, errRead
	}
	return bytes, nil
}
