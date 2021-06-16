package mvc

import (
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func ChatMsgList(db *Sql, value string) ([]ChatMsg, error) {
	var art []ChatMsg
	var result vo.ChatMsgListParams
	err := json.Unmarshal([]byte(value), &result)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", result)

	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(result.Token)
	if !b {
		return art, errors.New("token 失效")
	}
	sugar.Log.Info("claim := ", claim)
	sugar.Log.Error("Marshal data is  result := ", result)
	r := (result.PageNum - 1) * 3
	sugar.Log.Info("pageSize := ", result.PageSize)
	sugar.Log.Info("pageNum := ", result.PageNum)
	sugar.Log.Info("r := ", r)
	sugar.Log.Info("recordId ==== := ", result.RecordId)

	//这里 要修改   加上 where  参数 判断

	//todo
	//
	rows, err := db.DB.Query("SELECT * FROM chat_msg where record_id =? limit ?,?", result.RecordId, r, result.PageSize)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return art, errors.New("查询下载列表信息失败")
	}
	for rows.Next() {
		var dl ChatMsg
		err = rows.Scan(&dl.Id, &dl.ContentType, &dl.Content, &dl.FromId, &dl.ToId, &dl.Ptime, &dl.IsWithdraw, &dl.IsRead, &dl.RecordId)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return art, err
		}
		sugar.Log.Info("Query a entire data is ", dl)
		art = append(art, dl)
	}
	if err != nil {
		sugar.Log.Error("Insert into article  is Failed.", err)
		return art, err
	}
	sugar.Log.Info("Insert into article  is successful.")
	return art, nil

}
