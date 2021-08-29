package mvc

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func ChatSaveOfflineMsgs(db *Sql, value string) error {

	// 接收参数
	var msgs []vo.ChatSwapMsgParams

	sugar.Log.Debug("Request Param:", value)

	err := json.Unmarshal([]byte(value), &msgs)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return err
	}
	sugar.Log.Info("Marshal data is  ", msgs)

	// 会话是否存在
	chatMap := make(map[string]vo.ChatSwapMsgParams)
	for _, msg := range msgs {
		chatMap[msg.RecordId] = msg
	}

	for rid, msg := range chatMap {
		var recordId string
		err := db.DB.QueryRow("select id from chat_record where id = ?", rid).Scan(&recordId)
		if err != nil && err == sql.ErrNoRows {
			_, err := db.DB.Exec("INSERT INTO chat_record (id, name, from_id, to_id, ptime, last_msg) VALUES (?, ?, ?, ?, ?, ?)", msg.RecordId, "", msg.FromId, msg.ToId, msg.Ptime, "")
			if err != nil {
				sugar.Log.Error("INSERT INTO chat_record is Failed.", err)
				return err
			}
		}
	}

	// 保存消息
	for _, msg := range msgs {
		_, err := db.DB.Exec(
			"insert into chat_msg (id, content_type, content, from_id, to_id, ptime, is_with_draw, is_read, record_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
			msg.Id, msg.ContentType, msg.Content, msg.FromId, msg.ToId, msg.Ptime, msg.IsWithdraw, msg.IsRead, msg.RecordId)
		if err != nil {
			sugar.Log.Error("insert chat_msg is Failed.", err)
			continue
		}
	}

	// 更新最新消息
	for rid := range chatMap {
		var lastmsg string
		var isWithDraw int64
		err := db.DB.QueryRow("select is_with_draw, content from chat_msg where record_id = ? order by ptime desc", rid).Scan(&isWithDraw, &lastmsg)
		if err != nil {
			sugar.Log.Error("select chat_msg Failed.", err)
			return err
		}

		if isWithDraw == 1 {
			lastmsg = "撤回了一条消息"
		}

		_, err = db.DB.Exec("update chat_record set last_msg = ? WHERE id = ?", lastmsg, rid)
		if err != nil {
			return err
		}

		fmt.Printf("update chat_record set last_msg = %s WHERE id = %s", lastmsg, rid)
	}

	return nil
}
