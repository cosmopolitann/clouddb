package mvc

import (
	"encoding/json"
	"errors"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func ChatUserIdList(db *Sql, value string) (userIds []string, err error) {

	var req vo.ChatUserIdListParams

	sugar.Log.Debug("Request Param: ", value)
	err = json.Unmarshal([]byte(value), &req)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return userIds, err
	}
	sugar.Log.Info("Marshal data is  ", req)

	//token verify
	claim, b := jwt.JwtVeriyToken(req.Token)
	if !b {
		err = errors.New("jwt.JwtVeriyToken failed")
		return
	}

	sugar.Log.Info("claim := ", claim)

	userId := claim["id"].(string)

	// 查询会话列表
	rows, err := db.DB.Query("SELECT id, from_id, to_id, ptime, last_msg FROM chat_record WHERE from_id = ? OR to_id = ? ORDER BY ptime DESC", userId, userId)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
	}
	// 释放锁
	defer rows.Close()
	mapUserIds := make(map[string]string)
	for rows.Next() {
		var ri vo.ChatRecored
		err = rows.Scan(&ri.Id, &ri.FromId, &ri.ToId, &ri.Ptime, &ri.LastMsg)
		if err != nil {
			sugar.Log.Error("Query data is failed.Err is ", err)
			return
		}

		mapUserIds[ri.FromId] = ri.FromId
		mapUserIds[ri.ToId] = ri.ToId
	}

	for uid := range mapUserIds {
		userIds = append(userIds, uid)
	}

	return
}
