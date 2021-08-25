package cloud

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/cosmopolitann/clouddb/sugar"
)

//CopyFile
func TestCopyFile(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "/Users/apple/Desktop/xiaolong1.db")
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
"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjQxNDIwNzExNDIxNTQyODA5MCIsInBlZXJJZCI6IiIsIm5hbWUiOiLmtYvor5UiLCJwaG9uZSI6IiIsInNleCI6MCwibmlja25hbWUiOiLmtYvor5UiLCJpbWciOiIiLCJwdGltZSI6MTYyNzQ0NDAwOCwidXRpbWUiOjE2Mjc0NDQwMDgsInJvbGUiOiIxIiwiZXhwIjoxNjMwMDY5Nzg2fQ.8xl_c7JNPW9y3Xs62Y8Iu56CxwJTItwq67PvmFF6CKQ",
"parentId":"435120994852540416",
"ids":["435122292381454336"]
}`
	//b1, e := json.Marshal(fi)//
	//fmt.Println(ss)//
	//fmt.Println(b1)//
	resp := ss.CopyFile(string(value))
	fmt.Println("这是返回的数据 =", resp)

}
