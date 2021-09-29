package cloud

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	"testing"
)

func TestAddDelete(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "/Users/apple/Desktop/temp.db")
	if err != nil {
		panic(err)
	}
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ", d)
	e := d.Ping()
	fmt.Println(" Ping is failed,err:=", e)
	ss := Testdb(d)
	//插入数据
	var f1 = vo.CloudAddFolderParams{
		Id:       "123",
		FileName: "文学巨著",
		ParentId: "12",
		Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MDkzMzAyMDIxNjY5NTYwMzIiLCJleHAiOjE2MjU4ODk0NzZ9.OzEFVuB2FcRYurZiii1fpiAqX2KcesfS5arJfVJZQOI",
	}
	b1, e := json.Marshal(f1)
	fmt.Println(e)
	fmt.Println(string(b1))

	value := `{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjQxNDIwNzExNDIxNTQyODA5MCIsInBlZXJJZCI6IiIsIm5hbWUiOiLmtYvor5UiLCJwaG9uZSI6IiIsInNleCI6MCwibmlja25hbWUiOiLmtYvor5UiLCJpbWciOiIiLCJwdGltZSI6MTYyNzQ0NDAwOCwidXRpbWUiOjE2Mjc0NDQwMDgsInJvbGUiOiIxIiwiZXhwIjoxNjMzMTIyMTg2fQ.ABNWvSzav4-KP_nckC8C7ys6-1Fbj84LrqHK8OUHtbM","ids":["44","55"]}
`

	resp := ss.DeleteAll(value)
	fmt.Println("这是返回的数据 =", resp)
}
