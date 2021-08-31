package chat

import (
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	_ "github.com/mattn/go-sqlite3"
)

func TestChatSaveOfflineMsg(t *testing.T) {
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
	err = d.Ping()
	if err != nil {
		panic(err)
	}

	req := []vo.ChatSwapMsgParams{
		{
			Id:          "111",
			RecordId:    "416203557629337600_416418922095452160",
			ContentType: 2,
			Content:     "2222",
			FromId:      "416203557629337600",
			ToId:        "416418922095452160",
			Ptime:       1231233,
		},
		{
			Id:          "222",
			RecordId:    "416203556291354624_416418922095452160",
			ContentType: 2,
			Content:     "33333",
			FromId:      "416203556291354624",
			ToId:        "416418922095452160",
			Ptime:       1231234,
		},
	}
	value, _ := json.Marshal(req)

	ss := Testdb(d)

	resp := ss.ChatSaveOfflineMsgs(string(value))
	t.Log("获取返回的数据 :=  ", resp)
}
