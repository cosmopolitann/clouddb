package mvc

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"github.com/cosmopolitann/clouddb/vo"
	ipfsCore "github.com/ipfs/go-ipfs/core"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

func AddArticle(ipfsNode *ipfsCore.IpfsNode, db *Sql, value string, path string) error {
	sugar.Log.Info(" ----  AddArticle Method ----")
	sugar.Log.Info(" ----  Path :", path)
	var art vo.ArticleAddParams
	err := json.Unmarshal([]byte(value), &art)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err:", err)
		return errors.New(" Marshal article params is failed. ")
	}
	sugar.Log.Info("Marshal article params data : ", art)
	id := utils.SnowId()
	t := time.Now().Unix()
	stmt, err := db.DB.Prepare("INSERT INTO article (id,user_id,accesstory,accesstory_type,text,tag,ptime,play_num,share_num,title,thumbnail,file_name,file_size,external_href) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert into article table is failed.Err: ", err)
		return errors.New(" Insert into article table is failed. ")
	}
	sid := strconv.FormatInt(id, 10)
	//stmt.QueryRow()
	res, err := stmt.Exec(sid, art.UserId, art.Accesstory, art.AccesstoryType, art.Text, art.Tag, t, 0, 0, art.Title, art.Thumbnail, art.FileName, art.FileSize, art.ExternalHref)
	if err != nil {
		sugar.Log.Error(" Insert into article  is Failed.", err)
		return errors.New(" Execute query article table is failed. ")
	}
	l, _ := res.RowsAffected()
	if l == 0 {
		return errors.New(" Insert into article table is failed. ")
	}

	//--------------- publish msg ----------------
	// var ok bool
	topic := "/db-online-sync"
	var tp *pubsub.Topic
	ctx := context.Background()
	tp, ok := TopicJoin.Load(topic)
	if !ok {
		tp, err = ipfsNode.PubSub.Join(topic)
		if err != nil {
			sugar.Log.Error("PubSub.Join .Err is", err)
			return err
		}
		TopicJoin.Store(topic, tp)
	}
	sugar.Log.Info("Publish topic name :", "/db-online-sync")
	//step 1
	//query a article data
	var dl vo.ArticleResp
	err = db.DB.QueryRow("SELECT id,IFNULL(user_id,'null'),IFNULL(accesstory,'null'),IFNULL(accesstory_type,0),IFNULL(text,'null'),IFNULL(tag,'null'),IFNULL(ptime,0),IFNULL(play_num,0),IFNULL(share_num,0),IFNULL(title,'null'),IFNULL(thumbnail,'null'),IFNULL(file_name,'null'),IFNULL(file_size,0) from article where id=?;", sid).Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize)
	if err != nil && err != sql.ErrNoRows {
		sugar.Log.Error("Query article failed.Err is", err)
		return err
	}
	//

	//query user info.
	var dl1 vo.RespSysUser
	rows, err := db.DB.Query("select id,IFNULL(peer_id,'null'),IFNULL(name,'null'),IFNULL(phone,'null'),IFNULL(sex,0),IFNULL(ptime,0),IFNULL(utime,0),IFNULL(nickname,'null'),IFNULL(img,'null'),IFNULL(role,'2') from sys_user where id=?", art.UserId)
	if err != nil {
		sugar.Log.Error("AddUser Query data is failed.Err is ", err)
		return err
	}
	// ?????????
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&dl1.Id, &dl1.PeerId, &dl1.Name, &dl1.Phone, &dl1.Sex, &dl1.Ptime, &dl1.Utime, &dl1.NickName, &dl1.Img, &dl1.Role)
		if err != nil {
			sugar.Log.Error("AddUser Query scan data is failed.The err is ", err)
			return err
		}
		sugar.Log.Info(" AddUser Query a entire data is ", dl)
	}

	//the first step.
	var s3 UserAd
	s3.Type = "receiveUserRegister"
	s3.Data = dl1
	s3.FromId = ipfsNode.Identity.String()
	//marshal UserAd.
	//the second step
	sugar.Log.Info("--- second step ---")

	jsonBytes, err := json.Marshal(s3)
	if err != nil {
		sugar.Log.Error("Publish msg is failed.Err:", err)
		return err
	}
	sugar.Log.Info("Frwarding information:=", string(jsonBytes))
	sugar.Log.Info("Local PeerId :=", ipfsNode.Identity.String())
	//the  third  step .
	sugar.Log.Info("--- third step ---")

	err = tp.Publish(ctx, jsonBytes)
	if err != nil {
		sugar.Log.Error("Publish Err:", err)
		return err
	}

	//----
	var g PubSyncArticle
	g.Data = dl
	g.Type = "receiveArticleAdd"
	g.FromId = ipfsNode.Identity.String()
	//struct => json
	jsonBytes, err = json.Marshal(g)
	if err != nil {
		sugar.Log.Error("Marshal struct => json is failed.")
		return err
	}
	sugar.Log.Info("Forward the data to the public gateway.data:=", string(jsonBytes))

	err = tp.Publish(ctx, jsonBytes)
	if err != nil {
		sugar.Log.Error("Publish info failed.Err:", err)
		return err
	}
	sugar.Log.Info("---  Publish to other device  ---")
	//

	sugar.Log.Info("~~~~  Publish msg is successful.   ~~~~  ")

	// ????????????
	sugar.Log.Info("---  write sql to file.  ---")
	//
	f1, err1 := os.OpenFile(path+"update", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666) //????????????
	if err1 != nil {
		sugar.Log.Error(" Create update is failed. Err :", err1)
	}

	sugar.Log.Info("----- Local file is exist.  ----")

	// ??????????????? sql ??????

	sql := fmt.Sprintf("INSERT INTO article (id,user_id,accesstory,accesstory_type,text,tag,ptime,play_num,share_num,title,thumbnail,file_name,file_size,external_href) values ('%s','%s','%s',%d,'%s','%s',%d,%d,%d,'%s','%s','%s','%s','%s')\n", sid, art.UserId, art.Accesstory, art.AccesstoryType, art.Text, art.Tag, t, 0, 0, art.Title, art.Thumbnail, art.FileName, art.FileSize, art.ExternalHref)

	_, err = f1.WriteString(sql)
	if err != nil {
		sugar.Log.Error("-----  Write sql to update file is failed.Err:  ----", err)
	}
	sugar.Log.Info("-----sql :-----", sql)

	sugar.Log.Info("-----  Write sql to file is successful~~ ----")
	sugar.Log.Info(" ----  AddArticle Method  End ----")
	return nil
}

type PubSyncArticle struct {
	Type   string         `json:"type"`
	Data   vo.ArticleResp `json:"data"`
	FromId string         `json:"from"`
}

func ArticleAddTest(db *Sql, value string) error {
	sugar.Log.Info(" ----  AddArticle Method ----")
	var art vo.ArticleAddParams
	err := json.Unmarshal([]byte(value), &art)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err:", err)
		return errors.New(" Marshal article params is failed. ")
	}
	sugar.Log.Info("Marshal article params data : ", art)
	id := utils.SnowId()
	t := time.Now().Unix()
	stmt, err := db.DB.Prepare("INSERT INTO article values(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert into article table is failed.Err: ", err)
		return errors.New(" Insert into article table is failed. ")
	}
	sid := strconv.FormatInt(id, 10)
	//stmt.QueryRow()
	res, err := stmt.Exec(sid, art.UserId, art.Accesstory, art.AccesstoryType, art.Text, art.Tag, t, 0, 0, art.Title, art.Thumbnail, art.FileName, art.FileSize)
	if err != nil {
		sugar.Log.Error(" Insert into article  is Failed.", err)
		return errors.New(" Execute query article table is failed. ")
	}
	l, _ := res.RowsAffected()
	if l == 0 {
		return errors.New(" Insert into article table is failed. ")
	}
	return nil
}
