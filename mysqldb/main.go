package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
)

func InitDB(username string, password string, host string, port int, database string) (db *sql.DB, err error) {
	var portocol = strings.Join([]string{host, strconv.Itoa(port)}, ":")
	var connect = []string{username, ":", password, "@", "tcp(", portocol, ")", "/", database}
	DSN := strings.Join(connect, "")
	fmt.Println(DSN)
	db, err = sql.Open("mysql", DSN)
	return db, err
}

func main() {
	db, err := InitDB("root", "q13fvDipVup@dlfzsemq", "192.168.8.231", 3306, "ko_open")
	if err != nil {
		fmt.Printf("Connect err !! %v", err.Error())
		return
	}
	defer db.Close()
	db.Ping()
	resp, err := db.Query("show tables;")
	defer resp.Close()
	for resp.Next() {
		resp.Scan()
		fmt.Println()
	}
	if err != nil {
		fmt.Printf("query error %v", err.Error())
		return
	}
	fmt.Printf("%v")
}
