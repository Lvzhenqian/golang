package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)
var (
	engine *xorm.Engine
)

type User struct {
	Id 		int
	Name 	string
	Salt	string
	Age		int8
	Passwd	string
	Created	time.Time
	Update 	time.Time
}

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql","root:sd-9898w@tcp(127.0.0.1:3306)/psgo?charset=utf8")
	if err != nil{
		return}
	defer engine.Close()
	//engine.ShowSQL(true)
    c := User{1,"lv","cc",33,"abc",time.Now(),time.Now()}
	//errs := engine.Sync2(new(User))
	engine.Insert(&c)
	//fmt.Fprintf(os.Stderr,"err: %v",errs.Error())
}