package mvc

import (
	"encoding/json"
	"fmt"

	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

//更新用户信息
func ChatUsersUpdate(db *Sql, value string) error {
	var userlist []vo.ChatUserInfo

	sugar.Log.Debug("Request Param: ", value)
	err := json.Unmarshal([]byte(value), &userlist)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}

	for _, u := range userlist {
		_, err := db.DB.Exec("INSERT OR REPLACE INTO sys_user(id, peer_id, name, sex, nickname, img) VALUES (?, ?, ?, ?, ?, ?)",
			u.Id, u.PeerId, u.Name, u.Sex, u.Nickname, u.Img)
		if err != nil {
			return fmt.Errorf("update user err: %w", err)
		}
	}
	return nil
}
