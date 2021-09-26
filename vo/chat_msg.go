package vo

const (
	MSG_TYPE_WITHDRAW  = "receiveMsgWithDraw" // 消息：撤销聊天
	MSG_TYPE_NEW       = "receiveMsg"         // 消息：聊天消息
	MSG_TYPE_RECORD    = "receiveRecord"      // 消息：新会话
	MSG_TYPE_ACK       = "ack"                // 消息：消息回执
	MSG_TYPE_HEARTBEAT = "heartbeat"          // 消息：心跳检查

	CHAT_MSG_SWAP_TOPIC      = "xiaolong-chat-swap" // 消息接收主题：公共
	CHAT_MSG_SWAP_TOPIC_USER = "xlcs"               // 消息接收主题：个人
)

const (
	MSG_STATE_SENDING = 0  // 发送中
	MSG_STATE_SUCCESS = 1  // 发送成功
	MSG_STATE_FAILED  = -1 // 发送失败
)

type ChatListenHandler interface {
	HandlerChat(string)
}

type ChatFailMessageHandler interface {
	HandlerOfflineMessage(string)
}

type ChatPacketParams struct {
	Type    string      `json:"type"`    //类型
	Data    interface{} `json:"data"`    //数据
	From    string      `json:"from"`    //来源
	Receive string      `json:"receive"` //接受人
}

type ChatSendMsgParams struct {
	RecordId    string       `json:"recordId"`    //require     coment 消息记录id
	ContentType int64        `json:"contentType"` //       1 文本  2 表情 3 图片 4 文件
	Content     string       `json:"content"`     // require     coment 消息内容
	FromId      string       `json:"fromId"`      //require     coment 发送方id
	ToId        string       `json:"toId"`        //require     coment 对方id
	Token       string       `json:"token"`       //token
	Peer        ChatUserInfo `json:"peer"`        // peer
}

type ChatReadMsgParams struct {
	Ids   []string `json:"ids"`   //require
	Token string   `json:"token"` //token
}

type ChatSwapMsgParams struct {
	Id          string       `json:"id"`
	RecordId    string       `json:"recordId"`    //require     coment 消息记录id
	ContentType int64        `json:"contentType"` //       1 文本  2 表情 3 图片 4 文件
	Content     string       `json:"content"`     // require     coment 消息内容
	FromId      string       `json:"fromId"`      //require     coment 发送方id
	ToId        string       `json:"toId"`        //require     coment 对方id
	IsWithdraw  int64        `json:"isWithDraw"`  //require     coment 是否撤回         0 未撤回  1  撤回
	IsRead      int64        `json:"isRead"`      // require     coment 是否已读
	Ptime       int64        `json:"ptime"`
	Token       string       `json:"token"` //token
	User        ChatUserInfo `json:"user"`
}

type ChatAddRecordParams struct {
	Name   string `json:"name"`
	FromId string `json:"fromId"`
	ToId   string `json:"toId"`
	Token  string `json:"token"` //token
}

type ChatSwapRecordParams struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Img     string `json:"img"`
	FromId  string `json:"fromId"`
	ToId    string `json:"toId"`
	Ptime   int64  `json:"ptime"`
	LastMsg string `json:"lastMsg"`
	Token   string `json:"token"` //token
}

type ChatWithdrawMsgParams struct {
	MsgId  string `json:"id"`     //require     消息ID
	FromId string `json:"fromId"` //require     发送者ID
	ToId   string `json:"toId"`   //require     发送者ID
	Token  string `json:"token"`  //token
}

type ChatSwapWithdrawMsgParams struct {
	RecordId string `json:"recordId"` //require     coment 消息记录id
	MsgId    string `json:"id"`       //require     消息ID
	FromId   string `json:"fromId"`   //require     发送者ID
	ToId     string `json:"toId"`     //require     发送者ID
	Token    string `json:"token"`    //token
}

type ChatSwapAckParams struct {
	Type   string `json:"type"`   // receiveMsgWithDraw ｜ receiveMsg ｜ receiveRecord
	Id     string `json:"id"`     // 回执ID
	FromId string `json:"fromId"` // 发送者
	ToId   string `json:"toId"`   // 接收者
}

type UserBaseInfo struct {
	Id       string `json:"id"`       //id
	PeerId   string `json:"peerId"`   //peerId
	Name     string `json:"name"`     //名字
	Phone    string `json:"phone"`    //手机号
	Sex      int64  `json:"sex"`      //性别 0 未知 1 男 2 女
	Ptime    int64  `json:"ptime"`    //时间
	Nickname string `json:"nickname"` //昵称
	Img      string `json:"img"`      //图片
}

type OfflineMessage struct {
	Id          string `json:"id"`          //id
	RecordId    string `json:"recordId"`    //record_id
	FromId      string `json:"fromId"`      //from用户id
	ToId        string `json:"toId"`        //to用户id
	ContentType int64  `json:"contentType"` //附件类型
	Content     string `json:"content"`     //附件
	IsWithdraw  int64  `json:"isWithdraw"`  //是否撤回
	Ptime       int64  `json:"ptime"`       //时间
	IsRead      int64  `json:"isRead"`
}

type OfflineMessageV2 struct {
	User     UserBaseInfo     `json:"user"`     // 消息发送人
	Messages []OfflineMessage `json:"messages"` // 离线消息列表
}
