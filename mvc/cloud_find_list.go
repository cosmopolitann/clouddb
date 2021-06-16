package mvc

import (
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	"encoding/json"
	"errors"
)

//查询文件列表

func CloudFindList(db *Sql, value string)(data []File,e error) {
	var list vo.CloudFindListParams
	err := json.Unmarshal([]byte(value), &list)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", list)
	// 查询
	var arrfile []File
	rows, err := db.DB.Query("select * from cloud_file where user_id=? and parent_id=?", "",list.ParentId)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return arrfile,errors.New("查询下载列表信息失败")
	}
	for rows.Next() {
		var dl File
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.FileName, &dl.ParentId,&dl.Ptime, &dl.FileCid, &dl.FileSize, &dl.FileType,&dl.IsFolder)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return arrfile,err
		}
		sugar.Log.Info("Query a entire data is ", dl)
		arrfile=append(arrfile,dl)
	}
	sugar.Log.Info("Query all data is ", arrfile)

	sugar.Log.Info("FindList article  is successful.")

	return arrfile,nil

}
