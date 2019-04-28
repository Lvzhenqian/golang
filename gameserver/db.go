package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type SqlClient struct {
	db *sql.DB
}

func InitDB(username string, password string, host string, port int, database string, DbCharset string) (b *SqlClient) {
	var client = &SqlClient{}
	var err error
	protocol := fmt.Sprintf("%v:%v", host, port)
	DSN := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=%v&parseTime=true&loc=Local",
		username, password, protocol, database, DbCharset)
	if client.db, err = sql.Open("mysql", DSN); err != nil {
		panic(err.Error())
	} else {
		if PingErr := client.db.Ping(); PingErr != nil {
			panic(PingErr.Error())
		} else {
			return client
		}
	}
}

func (client *SqlClient) LoadGame() []string {
	rows, e := client.db.Query("SELECT `gameID`,`gameName`,`appKey`,`appSecret` FROM `battle_game`;")
	var ret []string
	if e != nil {
		panic(e.Error())
	}
	for rows.Next() {
		var gameID, gameName, appKey, appSecret string
		if err := rows.Scan(&gameID, &gameName, &appKey, &appSecret); err != nil {
			panic(err.Error())
		}
		ret = append(ret, []string{gameID, gameName, appKey, appSecret}...)
	}
	return ret
}

func (client *SqlClient) InsertGame(gameID, gameName, appKey, appSecret string) int64 {
	tx, e := client.db.Begin()
	if e != nil {
		panic(e.Error())
	}
	stmt, Perr := tx.Prepare("INSERT INTO `battle_game` (`gameID`,`gameName`,`appKey`,`appSecret`,`slide`) VALUE (?,?,?,?,?);")
	if Perr != nil {
		panic(Perr.Error())
	}
	if ret, err := stmt.Exec(gameID, gameName, appKey, appSecret, ""); err != nil {
		tx.Rollback()
	} else {
		ids, _ := ret.LastInsertId()
		return ids
	}
	if CommitErr := tx.Commit(); CommitErr != nil {
		tx.Rollback()
	}
	defer stmt.Close()
	return 0
}

func (client *SqlClient) DeleteGame(gameID string) int64 {
	stmt, e := client.db.Prepare("DELETE FROM battle_game WHERE gameID=?;")
	if e != nil {
		panic(e.Error())
	}
	defer stmt.Close()
	result, err := stmt.Exec(gameID)
	if err != nil {
		panic(err)
	}
	ids,_ := result.LastInsertId()
	return ids
}
