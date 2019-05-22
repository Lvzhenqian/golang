package main

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	result := bson.M{}
	if err := session.DB("admin").Run(bson.D{{"serverStatus", 1}}, &result); err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}
	if err := session.DB("test").Run("dbstats", &result); err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}
	if err := session.DB("admin").Run("replSetGetStatus", &result); err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}
}
