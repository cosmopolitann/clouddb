package chat

import (
	"database/sql"
	"testing"

	"github.com/cosmopolitann/clouddb/sugar"
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

	// req := []vo.ChatSwapMsgParams{
	// 	{
	// 		Id:          "111",
	// 		RecordId:    "416203557629337600_416418922095452160",
	// 		ContentType: 2,
	// 		Content:     "2222",
	// 		FromId:      "416203557629337600",
	// 		ToId:        "416418922095452160",
	// 		Ptime:       1231233,
	// 	},
	// 	{
	// 		Id:          "222",
	// 		RecordId:    "416203556291354624_416418922095452160",
	// 		ContentType: 2,
	// 		Content:     "33333",
	// 		FromId:      "416203556291354624",
	// 		ToId:        "416418922095452160",
	// 		Ptime:       1231234,
	// 	},
	// }
	// value, _ := json.Marshal(req)

	sval := `[{"user":{"id":"438760061880242176","peerId":"Qmc9gVATXWEdXPhep4iSMXTwcodt6JVYyf4vM4tNWbt5Pn","name":"","phone":"","sex":1,"ptime":1630314083,"nickname":"dragon小龙虾小龙虾","img":"bafybeieqyv6ddx4flbjfuqxvw5w5smrxijgs5xjjyju4dz4xgubngzmb7e"},"messages":[{"id":"439033666128056321","recordId":"436205633679659008_438760061880242176","fromId":"438760061880242176","toId":"436205633679659008","contentType":1,"content":"{\"text\":\"33\",\"avatar\":\"QmXiRuuprLwAPdpZsKCBmpMcMPmFd67J31ov5iSshoDT2E\",\"name\":\"郑亚利\"}","isWithdraw":0,"ptime":1630379316}]},{"user":{"id":"439020064117624832","peerId":"Qmc9gVATXWEdXPhep4iSMXTwcodt6JVYyf4vM4tNWbt5Pn","name":"","phone":"","sex":1,"ptime":1630376072,"nickname":"dragon24832","img":"bafkreif7zaxv7vncryi4aqhq3zkb36ykh6brq5luxiaaflfezqjy3gjsnq"},"messages":[{"id":"439033666128056332","recordId":"436205633679659008_439020064117624832","fromId":"439020064117624832","toId":"436205633679659008","contentType":2,"content":"{\"text\":\"122121212\",\"avatar\":\"QmXiRuuprLwAPdpZsKCBmpMcMPmFd67J31ov5iSshoDT2E\",\"name\":\"郑亚利\"}","isWithdraw":0,"ptime":1630379319}]}]`

	ss := Testdb(d)

	resp := ss.ChatSaveOfflineMsgsV2(sval)
	t.Log("获取返回的数据 :=  ", resp)
}
