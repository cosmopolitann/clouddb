package chat

import (
	"database/sql"
	"fmt"
	"github.com/cosmopolitann/clouddb/mvc"
	"github.com/cosmopolitann/clouddb/sugar"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestChatMsg(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "/Users/apple/winter/D-cloud/tables/foo.db")
	if err!=nil{
		panic(err)
	}
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ",d)
	e:=d.Ping()
	fmt.Println(" Ping is failed,err:=",e)
	ss:= Testdb(d)
	value:=`{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MDkzMzAyMDIxNjY5NTYwMzIiLCJleHAiOjE2MjU4ODk0NzZ9.OzEFVuB2FcRYurZiii1fpiAqX2KcesfS5arJfVJZQOI","id":"4324","contentType":2,"content":"20","fromId":"410199041234702336","toId":"409330202166956032","isWithdraw":2,"isRead":3,"recordId":"110"}
`
	resp:= ss.ChatAddMsg(value)
	t.Log("获取返回的数据 :=  ",resp)

}
func Testdb(sq *sql.DB) mvc.Sql {
	return mvc.Sql{DB: sq}
}