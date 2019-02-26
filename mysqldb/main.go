package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
type DbClient struct {
	b *sql.DB
}

func InitDB(username string, password string, host string, port int, database string, DbCharset string) (b *DbClient) {
	var db = &DbClient{}
	var err error
	protocol := fmt.Sprintf("%v:%v", host, port)
	DSN := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=%v&parseTime=true&loc=Local",
		username, password, protocol, database, DbCharset)
	if db.b, err = sql.Open("mysql", DSN); err != nil {
		panic(err.Error())
	} else {
		if PingErr := db.b.Ping(); PingErr != nil {
			panic(PingErr.Error())
		} else {
			return db
		}
	}
}

func (db *DbClient) InsertToDB(sq string,v ...interface{}) (int64,error) {
	if stmt, e := db.b.Prepare(sq); e != nil{
		panic(e.Error())
	} else {
		result, err := stmt.Exec(v...)
		defer stmt.Close()
		if err != nil {
			panic(err.Error())
		}
		ids,GidErr := result.LastInsertId()
		return ids,GidErr
	}
}


func main() {
	db := InitDB("root", "q13fvDipVup@dlfzsemq",
		"192.168.8.20", 13306, "ko_open","utf8")
	defer db.b.Close()
	ids, e := db.InsertToDB("insert into `ko_open`.`battle_game` (`gameID`,`appKey`,`appSecret`,`gameName`,`slide`) value (?,?,?,?,?)",
		"201226",
		"b90b138ab59f4e289fbd58182d4187bb",
		"bf53a5a8b29c4198abe0e71c9e05d465",
		"frigate", "")
	if e != nil{
		panic(e.Error())
	}
	fmt.Println(ids)
}
