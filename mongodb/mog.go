package main

import (
	"flag"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"strconv"
	"strings"
)

var result map[string]interface{}

// 获取mongo 服务器指标
func mongo(ip string, port int) {
	url := fmt.Sprintf("mongodb://%s:%d/",ip,port)
	session, SessionErr := mgo.Dial(url)
	if SessionErr != nil {
		panic(SessionErr)
	}
	defer session.Clone()
	session.SetMode(mgo.Monotonic,true)
	resultErr := session.DB("admin").Run(bson.D{
		{"serverStatus",1},
	},&result)
	if resultErr != nil {
		panic(resultErr)
	}
	return
}

// 递归获取json数据
func PrintMapInterface(key []string ,values map[string]interface{}) {
	switch v :=values[key[0]].(type) {
	case map[string]interface{}:
		secondkey := key[1:]
		if len(secondkey) > 0 {
			// 递归调用类型判断
			PrintMapInterface(secondkey,v)
		} else {
			fmt.Printf("%v\n",v)
		}
	default:
		fmt.Println(v)
	}
}

func main() {
	flag.Parse()
	args := flag.Args()
	address := strings.Split(args[0],":")
	port ,PortErr:= strconv.Atoi(address[1])
	if PortErr != nil{
		panic(PortErr)
	}
	mongo(address[0],port)
	cmd := args[1:]
	PrintMapInterface(cmd,result)
}