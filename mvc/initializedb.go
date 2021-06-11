package mvc

import (
	"github.com/cosmopolitann/clouddb/sugar"
	"database/sql"
)
type Sql struct {
	DB *sql.DB

}



func Newdb(path string) Sql {
	return Sql{DB: InitDB(path)}
}

func InitDB(path string)(*sql.DB){
	//
	//mvc, err := sql.Open("sqlite3", path)
	sugar.Log.Info("Start Open Sqlite3 Database.")
	db, err := sql.Open("sqlite3", "/Users/apple/winter/D-cloud/tables/foo.db")
	checkErr(err)
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ",db)

	return db
}
func checkErr(err error) {
	if err != nil {
		sugar.Log.Error("The connection to the database failed.")
		panic(err)
	}
}