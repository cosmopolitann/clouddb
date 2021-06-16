package mvc

import (
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

//查询文件列表

func CloudFolderList(db *Sql, value string) (data []File, e error) {
	var arrfile []File

	var list vo.CloudFolderListParams
	err := json.Unmarshal([]byte(value), &list)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", list)
	//验证 token 是否满足

	claim, b := jwt.JwtVeriyToken(list.Token)
	if !b {
		return arrfile, errors.New("token 验证不通过")
	}
	sugar.Log.Info("claim := ", claim)

	// 查询
	rows, err := db.DB.Query("select * from cloud_file where parent_id=? and is_folder=? and user_id=?", list.ParentId, 1, claim["UserId"].(string))
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return arrfile, errors.New("查询下载列表信息失败")
	}
	for rows.Next() {
		var dl File
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.FileName, &dl.ParentId, &dl.Ptime, &dl.FileCid, &dl.FileSize, &dl.FileType, &dl.IsFolder)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return arrfile, err
		}
		sugar.Log.Info("Query a entire data is ", dl)
		arrfile = append(arrfile, dl)
	}
	sugar.Log.Info("Query all data is ", arrfile)
	return arrfile, nil

	sugar.Log.Info("Insert into article  is successful.")

	return arrfile, nil

}
