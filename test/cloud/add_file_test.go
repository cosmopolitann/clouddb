package cloud

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cosmopolitann/clouddb/mvc"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestAddFile(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "/Users/apple/winter/D-cloud/tables/foo.db")
	if err != nil {
		panic(err)
	}
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ", d)
	e := d.Ping()
	fmt.Println(" Ping is failed,err:=", e)
	ss := Testdb(d)
	//插入数据
	var fi = vo.CloudAddFileParams{
		Id:       "409054047794892800",
		FileName: "我爱中国",
		ParentId: "0",
		FileCid:  "Qmcid12312",
		FileSize: 100,
		FileType: 0,
		Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MTExOTI3MjA0NDc1MDg0ODAiLCJleHAiOjE2MjYzMzM1NzJ9.EDv7k__JEugp-57RU3wDsJ1Swa3C0t-Ofr4KeQQhzeA",
	}

	b1, e := json.Marshal(fi)
	fmt.Println(e)
	fmt.Println("这是 json 数据", string(b1))

	resp := ss.AddFile(string(b1))
	fmt.Println("这是返回的数据 =", resp)

}
func Testdb(sq *sql.DB) mvc.Sql {
	return mvc.Sql{DB: sq}
}
