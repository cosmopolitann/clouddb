package chat

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	_ "github.com/mattn/go-sqlite3"
)

func TestChatUsersUpdate(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "/Users/apple/Projects/clouddb/tables/xiaolong.db")
	if err != nil {
		panic(err)
	}
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ", d)
	e := d.Ping()
	fmt.Println(" Ping is failed,err:=", e)
	ss := Testdb(d)

	// token, _ := jwt.GenerateToken("416203557629337600", "peerid", "name", "phone", "nickname", "img", "2", 0, 1, 1, 30*24*60*60)

	req := []vo.ChatUserInfo{
		{
			Id:       "123",
			PeerId:   "peerid",
			Name:     "name",
			Sex:      1,
			Nickname: "nickname",
			Img:      "imge",
		},
		{
			Id:       "12345",
			PeerId:   "peerid222",
			Name:     "name222",
			Sex:      1,
			Nickname: "nickname22",
			Img:      "imge22",
		},
	}

	value, _ := json.Marshal(req)

	resp := ss.ChatUsersUpdate(string(value))
	fmt.Println("获取返回的数据 :=  ", resp)

}
