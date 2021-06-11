package cloud

import (
	"github.com/cosmopolitann/clouddb/sugar"
	"database/sql"
	"fmt"
	"testing"
)

//CopyFile
func TestMoveFile(t *testing.T) {
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
	value := `{
"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MDkzMzAyMDIxNjY5NTYwMzIiLCJleHAiOjE2MjU4ODk0NzZ9.OzEFVuB2FcRYurZiii1fpiAqX2KcesfS5arJfVJZQOI",
    "parentId":"1234",
    "ids":["409330202166956032"]
}`
	//b1, e := json.Marshal(fi)
	//fmt.Println(ss)
	//fmt.Println(b1)
	resp := ss.MoveFile(string(value))
	fmt.Println("这是返回的数据 =", resp)

}
