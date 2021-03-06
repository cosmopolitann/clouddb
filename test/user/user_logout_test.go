package user

import (
	"github.com/cosmopolitann/clouddb/sugar"

	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestUserDel(t *testing.T) {
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
	value := `{"id":"409283588031254528","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiIxODg4IiwiZXhwIjoxNjI1ODc4MzYxfQ.FkABVsrygfxKq2_GWP5pG2G9oYpUqD1yw5cA3boB-Dc"}
`
	resp := ss.UserLoginOut(value)
	t.Log("获取返回的数据 :=  ", resp)

}
