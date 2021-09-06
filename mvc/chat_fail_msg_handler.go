package mvc

import (
	"github.com/cosmopolitann/clouddb/vo"
)

var FailMsgHandler vo.ChatFailMessageHandler

// 设置失败消息回调
func ChatFailMsgHandler(omh vo.ChatFailMessageHandler) error {

	FailMsgHandler = omh

	return nil
}
