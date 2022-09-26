package bot

import "fmt"

// receive
const (
	// 信息种类
	MsgPostType = "message"
	ReqPostType = "request"
	NtsPostType = "notice"
	MtaPostType = "meta_event"

	PriMsgType = "private"
	GruMsgType = "group"

	FriSubMsgType  = "friend"
	GruSubMsgType  = "group"
	GrusSubMsgType = "group_self"
	NorSubMsgType  = "normal"
	AoySubMsgType  = "anonymous" // 匿名
	NtsSubMsgType  = "notice"    // 群内禁止私聊的通告

	FriRequestType = "friend"
	GruRequestType = "group"

	GruUplNtsType = "group_upload"
	GruAdmNtsType = "group_admin"
	GruDecNtsType = "group_decrease"
	GruIncNtsType = "group_increase"
	GruBanNtsType = "group_ban"
	GruRecNtsType = "group_recall"
	GruCadNtsType = "group_card"
	FriAddNtsType = "friend_add"
	FriRecNtsType = "friend_recall"
	OffLinNtsType = "offline_file"
	CliStuNtsType = "client_status"
	EssNtsType    = "essence"
	NotifyNtsType = "notify"

	HonNotifySubType = "honor"
	PokNotifySubType = "poke"
	lukNotifySubType = "lucky_king"

	OwnRoleType = "owner"
	AdmRoleType = "admin"
	MebRoleType = "member"

	lifeMetaType = "lifecycle"
	hertMetaType = "heartbeat"
)

type RecKind struct {
	PostType string `json:"post_type"`
}

type RecMessage struct {
	Sender      Asender `json:"sender"`
	SubType     string  `json:"sub_type"`
	TempSource  string  `json:"temp_source"`
	MessageType string  `json:"message_type"`
	GroupId     int64   `json:"group_id"`
	Message     string  `json:"message"`
}

func (msg *RecMessage) String() string {
	if msg.MessageType == GruMsgType {
		return fmt.Sprintf("来自群: %d，用户: %s，的消息: \"%s\"",
			msg.GroupId,
			msg.Sender.String(),
			msg.Message)
	}
	return fmt.Sprintf("来自用户: %s，内容: \"%s\"",
		msg.Sender.String(),
		msg.Message)
}

type RecRequest struct {
	RequestType string `json:"request_type"`
	GroupId     int64  `json:"group_id"`
	UserId      int64  `json:"user_id"`
	Flag        string `json:"flag"`
	SubType     string `json:"sub_type"`
}

func (req *RecRequest) String() string {
	switch req.RequestType {
	case FriRequestType:
		return fmt.Sprintf("[好友邀请: %d]", req.UserId)
	case GruRequestType:
		return fmt.Sprintf("[群邀请: %d，来自: %d]", req.GroupId, req.UserId)
	default:
		return fmt.Sprintln("奇怪的邀请: ", *req)
	}
}

type RecNotice struct {
	NoticeType string `json:"notice_type"`
	SubType    string `json:"sub_type"`
	SenderId   int64  `json:"sender_id"`
	TargetId   int64  `json:"target_id"`
	GroupId    int64  `json:"group_id"`
}

func (nts *RecNotice) String() string {
	if nts.GroupId != 0 {
		return fmt.Sprintf("[消息: %s %s, 发送ID: %d，目标ID: %d, 群: %d]",
			nts.NoticeType,
			nts.SubType,
			nts.SenderId,
			nts.TargetId,
			nts.GroupId)
	}
	return fmt.Sprintf("[消息: %s %s, 发送ID: %d，目标ID: %d]",
		nts.NoticeType,
		nts.SubType,
		nts.SenderId,
		nts.TargetId)
}

type RecMeta struct {
	MetaEvenType string `json:"meta_event_type"`
	SubType      string `json:"sub_type"`
}

func (meta *RecMeta) String() string {
	return fmt.Sprintf("[元事件: %s, 类型: %s]",
		meta.MetaEvenType,
		meta.SubType)
}

type Asender struct {
	// base
	UserId   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
	// group
	Cardname string `json:"card"`
	Area     string `json:"area"`
	Level    string `json:"level"`
	Role     string `json:"role"`
	Title    string `json:"title"`
}

func (sd *Asender) String() string {
	if sd.Cardname != "" {
		return fmt.Sprintf("[QQ: %d, 群名片: %s, 职位: %s]",
			sd.UserId,
			sd.Cardname,
			sd.Role)
	}
	return fmt.Sprintf("[QQ: %d, 名称: %s]", sd.UserId, sd.Nickname)
}

// send thing
const (
	echo = "good"

	SendMsgApi    = "send_msg"
	DeleteMsgApi  = "delete_msg"
	GetImgApi     = "get_image"
	GroupKickApi  = "set_group_kick"
	GroupBanApi   = "set_group_ban"
	GroupCardApi  = "set_group_card"
	GroupLeaveApi = "set_group_leave"

	FriRequestApi      = "set_friend_add_request"
	GruRequestApi      = "set_group_add_request"
	GetStrangerInfoApi = "get_stranger_info"
	CanSendImgApi      = "can_send_image"
)

type SendRespondJson struct {
	Action string      `json:"action"`
	Params interface{} `json:"params"`
	Echo   string      `json:"echo"`
}
