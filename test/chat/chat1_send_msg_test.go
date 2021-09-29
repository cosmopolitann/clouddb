package chat

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/test/myipfs"
	"github.com/cosmopolitann/clouddb/vo"
	_ "github.com/mattn/go-sqlite3"
)

func TestChatSendMsg(t *testing.T) {
	// f, err := os.OpenFile("./mem.prof", os.O_RDWR|os.O_CREATE, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// cpuf, err := os.OpenFile("./cpu.prof", os.O_RDWR|os.O_CREATE, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer cpuf.Close()
	// pprof.StartCPUProfile(cpuf)

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

	token, _ := jwt.GenerateToken("416419138592841728", "peerid", "name", "phone", "nickname", "img", "2", 0, 1, 1, 30*24*60*60)

	req := vo.ChatSendMsgParams{
		RecordId:    "416419138592841728_449268758863679488",
		ContentType: 2,
		Content:     "content 222222223333",
		FromId:      "416419138592841728",
		ToId:        "449268758863679488",
		Token:       token,
		Peer: vo.ChatUserInfo{
			Id:       "449268758863679488",
			PeerId:   "xxxxyyyyy",
			Name:     "ooxxyy",
			Sex:      1,
			Nickname: "nickname222",
			Img:      "imgabcdxxx",
		},
	}
	value, _ := json.Marshal(req)

	ss := Testdb(d)

	node, err := myipfs.GetIpfsNode("/Users/apple/projects/clouddb/test/chat/.ipfs")
	if err != nil {
		sugar.Log.Info("xxxxx----", err)
		panic(err)
	}

	// // h2ID, _ := peer.Decode("12D3KooWS8qWyGimuUgDjakUFGJkDgvGYcMEjnj5xqojeDwf1rZm")
	// h2ID, _ := peer.Decode("12D3KooWMUCCUigkLYryEJpGC1DdnJV87x8GozccreW2SVgK7KXW")

	// addr, err := node.DHT.FindPeer(context.Background(), h2ID)
	// if err != nil {
	// 	fmt.Println("find peer err:", err)
	// }

	// fmt.Println("addr:", addr)

	var cl ChatFailMessageHandler

	resp := ss.ChatSendMsg(node, string(value), &cl)
	t.Log("获取返回的数据 :=  ", resp)

	// for i := 0; i < 5; i++ {
	// 	time.Sleep(10 * time.Second)
	// }

	// pprof.StopCPUProfile()
	// pprof.WriteHeapProfile(f)
}

type ChatFailMessageHandler struct{}

func (cl *ChatFailMessageHandler) HandlerOfflineMessage(abc string) {
	fmt.Println("TestChatSendMsg----\n", abc, "3333-----")
}
