package mvc

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func ChatSaveOfflineMsgsV2(db *Sql, value string) error {
	// 接收参数
	var params []vo.OfflineMessageV2

	sugar.Log.Debug("Request Param:", value)

	err := json.Unmarshal([]byte(value), &params)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return err
	}
	sugar.Log.Info("Marshal data is  ", params)

	for _, param := range params {

		user := param.User
		msgs := param.Messages

		// 更新用户信息
		if user.Id != "" {
			_, err = db.DB.Exec("INSERT OR REPLACE INTO sys_user(id, peer_id, name, phone, sex, nickname, img) VALUES (?, ?, ?, ?, ?, ?, ?)",
				user.Id, user.PeerId, user.Name, user.Phone, user.Sex, user.Nickname, user.Img)
			if err != nil {
				return err
			}
		}

		// 会话是否存在
		chatMap := make(map[string]vo.OfflineMessage)
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
				lastmsg = `{"text":"撤回了一条消息"}`
			}

			_, err = db.DB.Exec("update chat_record set last_msg = ? WHERE id = ?", lastmsg, rid)
			if err != nil {
				return err
			}

			fmt.Printf("update chat_record set last_msg = %s WHERE id = %s", lastmsg, rid)
		}

	}

	return nil
}

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
			lastmsg = `{"text":"撤回了一条消息"}`
		}

		_, err = db.DB.Exec("update chat_record set last_msg = ? WHERE id = ?", lastmsg, rid)
		if err != nil {
			return err
		}

		fmt.Printf("update chat_record set last_msg = %s WHERE id = %s", lastmsg, rid)
	}

	return nil
}

func ChatGetOfflineMsgCount(db *Sql, value string) int {
	var num int
	var msg vo.OfflineMessageCount

	err := json.Unmarshal([]byte(value), &msg)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return num
	}

	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(msg.Token)
	if !b {
		sugar.Log.Error("token 失效")
		return 0
	}

	userId := claim["id"].(string)

	err = db.DB.QueryRow("select count(*) as num from chat_msg where is_read = 0 and to_id = ?", userId).Scan(&num)
	if err != nil {
		sugar.Log.Error("select chat_msg Failed.", err)
		return 0
	}

	return num
}
