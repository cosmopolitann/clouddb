package sync

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/cosmopolitann/clouddb/sugar"
)

func TestQueryAllData(t *testing.T) {
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
	value := `{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjQxNDIwNzExNDIxNTQyODA5MCIsInBlZXJJZCI6IiIsIm5hbWUiOiLmtYvor5UiLCJwaG9uZSI6IiIsInNleCI6MCwibmlja25hbWUiOiLmtYvor5UiLCJpbWciOiIiLCJwdGltZSI6MTYyNzQ0NDAwOCwidXRpbWUiOjE2Mjc0NDQwMDgsInJvbGUiOiIxIiwiZXhwIjoxNjMzMTIyMTg2fQ.ABNWvSzav4-KP_nckC8C7ys6-1Fbj84LrqHK8OUHtbM"}`
	path := "/Users/apple/winter/offline/"

	resp := ss.SyncQueryAllData(value, path)
	fmt.Println("这是返回的数据 =", resp)

}
func TestDatabaseMigration(t *testing.T) {
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
	token := `{"Token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjQxNDIwNzExNDIxNTQyODA5MCIsInBlZXJJZCI6IiIsIm5hbWUiOiLmtYvor5UiLCJwaG9uZSI6IiIsInNleCI6MCwibmlja25hbWUiOiLmtYvor5UiLCJpbWciOiIiLCJwdGltZSI6MTYyNzQ0NDAwOCwidXRpbWUiOjE2Mjc0NDQwMDgsInJvbGUiOiIxIiwiZXhwIjoxNjMzMTIyMTg2fQ.ABNWvSzav4-KP_nckC8C7ys6-1Fbj84LrqHK8OUHtbM"}`
	resp := ss.SyncDatabaseMigration(token, "/Users/apple/winter/offline/", "QmPUwX3VzLdxTy4c3f4PVp6SxBaYgoVZEQCYoCNSTDbk1o")
	fmt.Println("这是返回的数据 =", resp)
}

