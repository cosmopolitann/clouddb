package mvc

import (
	"errors"
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"github.com/cosmopolitann/clouddb/vo"
	"github.com/goinggo/mapstructure"
	"strconv"
	"strings"
	"time"
)

func AddFolder(db *Sql, value string) error {
	//add folder
	sugar.Log.Info("---- 开始执行 新建文件夹 方法 ---- ")

	var f vo.CloudAddFolderParams
	f1 := ConvertString(value, f)
	err := mapstructure.Decode(f1, &f)


	sugar.Log.Info("---- 开始 解析 参数  ---- 参数：",f)

	sugar.Log.Info("Decode data  is  ", f)
	sugar.Log.Info("参数 ParentId: ", f.ParentId)
	sugar.Log.Info("参数 FileName:", f.FileName)
	sugar.Log.Info("参数 Id:", f.Id)
	sugar.Log.Info("参数 Token:", f.Token)


	if err != nil {
		sugar.Log.Error("Decode is failed.", err)
		return errors.New("解析map失败")
	}
	//token
	claim, b := jwt.JwtVeriyToken(f.Token)
	if !b {
		return err
	}
	sugar.Log.Info("校验 token 成功   Token:", claim["UserId"])

	sugar.Log.Info("claim := ", claim)

	userId := claim["UserId"].(string)

	sugar.Log.Info("  查看 user id  ",userId)

	sugar.Log.Info("判断文件夹 是否 不满足 格式 成功 1")

	e := IsFormat(f)

	if e != nil {
		return err
	}
	sugar.Log.Info("判断文件夹 是否 不满足 格式 成功  2")

	//count,in:= InsertIntoData(db,f,userId.(string))
	//	if in!=nil || count==0{
	//		return errors.New("创建文件夹失败")
	//	}

	sugar.Log.Info("-- 查找是否有 相同名字的文件夹 ---")
	c, _ := FindOneDirIsExist(db, f)
	sugar.Log.Info("--  c == ---",c)

	if c == 0 {
		count, in := InsertIntoData(db, f, userId)
		if in != nil || count == 0 {
			return errors.New("创建文件夹失败")
		}
	}

	//-1 代表 文件，名字相同 但 cid 不相同
	if c == 1 {
		//后缀名
		timeUnix := time.Now().Unix()
		timeUnixStr := strconv.FormatInt(timeUnix, 10)
		f.FileName = f.FileName + "_" + timeUnixStr
		count, in := InsertIntoData(db, f, userId)
		if in != nil || count == 0 {
			return errors.New("创建文件夹失败")
		}
	}

	//stmt, err := db.DB.Prepare("INSERT INTO cloud_file values(?,?,?,?,?,?,?,?,?,?)")
	//if err != nil {
	//	sugar.Log.Error("Insert into cloud_file table is failed.",err)
	//	return err
	//}
	//sid := strconv.FormatInt(id, 10)
	//res, err := stmt.Exec(sid,f.UserId ,f.FileName, f.ParentId,t ,f.FileCid,f.FileSize,f.FileStatus,f.FileType,f.IsFolder)
	//if err != nil {
	//	sugar.Log.Error("Insert into file  is Failed.",err)
	//	return err
	//}
	//sugar.Log.Info("Insert into file  is successful.")
	//l,_:=res.RowsAffected()
	////
	//fmt.Println(" l =",l)
	//先查询一下本层有没有相同文件名，否则不能创建文件夹

	//the create folder isexist.if exist,add local timestamp suffix
	sugar.Log.Info("--  新建 文件夹 成功  返回  ---")

	return nil
}
func InsertIntoData(db *Sql, f vo.CloudAddFolderParams, userId string) (c int64, e error) {
	//insert into
	//snowId
	id := utils.SnowId()
	//t:=time.Now().Format("2006-01-02 15:04:05")
	t := time.Now().Unix()
	sugar.Log.Info("--  开始  插入文件夹  ---")
	sugar.Log.Info("--  参数信息   ParentId---",f.ParentId)
	sugar.Log.Info("--  参数信息   ParentId---",f.FileName)
	sugar.Log.Info("--  参数信息   ---",f)

	stmt, err := db.DB.Prepare("INSERT INTO cloud_file values(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert into cloud_file table is failed.", err)
		return 0, err
	}
	sid := strconv.FormatInt(id, 10)

	res, err := stmt.Exec(sid, userId, f.FileName, f.ParentId, t, "", 0, 0, 1)
	if err != nil {
		sugar.Log.Error("Insert into file  is Failed.", err)
		return 0, err
	}
	sugar.Log.Info("Insert into file  is successful.")
	l, err := res.RowsAffected()
	if l==0{
		return 0,err
	}

	sugar.Log.Info("--  插入文件的 雪花id  = ---",sid)

	return l, nil
}

func IsFormat(f vo.CloudAddFolderParams) error {
	sugar.Log.Info("判断文件夹 是否 不满足 格式 成功 3")

	pId, err := strconv.Atoi(f.ParentId)
	if err != nil {
		return err
	}
	if pId < 0 {
		return errors.New("参数不能为负数")
	}
	sugar.Log.Info("判断文件夹 是否 不满足 格式 成功 4")

	if IsEmptyRename(f.FileName) {
		return errors.New("文件夹名称不能为空")
	}
	if IsLenRename(f.FileName) {
		return errors.New("文件夹名称过长")
	}
	if f.FileName[0] == vo.IllegalPoint {
		return errors.New("文件夹不能为包含非法字符")
	}
	if find := strings.Contains(f.FileName, "/"); find {
		return errors.New("文件夹不能为包含非法字符")
	}
	if find := strings.Contains(f.FileName, "\\"); find {
		return errors.New("文件夹不能为包含非法字符")
	}
	if find := strings.Contains(f.FileName, "*"); find {
		return errors.New("文件夹不能为包含非法字符")
	}
	sugar.Log.Info("判断文件夹 是否 不满足 格式 成功 5")

	return nil
}

func IsLenRename(rename string) bool {
	if len(rename) > 256 {
		return true
	}
	return false
}
func IsEmptyRename(rename string) bool {
	if len(rename) == 0 {
		return true
	}
	return false
}

func FindOneDirIsExist(mvc *Sql, d vo.CloudAddFolderParams) (int64, error) {
	//查询数据
	var f File
	rows, _ := mvc.DB.Query("SELECT * FROM cloud_file where file_name=? and parent_id=? and is_folder=?", d.FileName, d.ParentId, 1)
	for rows.Next() {
		err := rows.Scan(&f.Id, &f.UserId, &f.FileName, &f.ParentId, &f.Ptime, &f.FileCid, &f.FileSize, &f.FileType, &f.IsFolder)
		if err != nil {
			return 0, err
		}
	}

	if f.Id != "" {
		//如果文件 cid 不相等  返回 0
		return 1, nil
	}

	return 0, nil
}
