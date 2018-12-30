package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
	"time"
)
type User struct {
	Id int64
	Name string
	Salt string
	Age int
	Passwd string `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}


func CreateTable() {
	engine, err := xorm.NewEngine("mysql","root:sd-9898w@tcp(192.168.8.12:3306)/gopub?charset=utf8")
	if err != nil{
		return
	}
	errs := engine.Sync2(new(User))
	if errs != nil{
		fmt.Fprintf(os.Stderr,"err: %v\n",errs.Error())
		return
	}
}

func main() {
	CreateTable()
	//route := gin.Default()
	//route.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200,gin.H{"message":"pong"})
	//})
	//route.POST("/post", func(context *gin.Context) {
	//	var json conf.PythonConf
	//	if err := context.ShouldBindJSON(&json); err !=nil{
	//		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	//	}
	//	context.JSON(http.StatusOK,gin.H{"status":"successfuly"})
	//	fmt.Printf("%v",json)
	//})
	//route.Run()
}