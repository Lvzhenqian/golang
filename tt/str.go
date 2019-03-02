package main

import (
	"fmt"
	"github.com/knocknote/vitess-sqlparser/sqlparser"
)

func main() {
	stmt, err := sqlparser.Parse("select * from user_items where user_id=1 order by created_at limit 3 offset 10")
	if err != nil {
		panic(err)
	}
	sqlparser.Walk(func(node sqlparser.SQLNode) (kontinue bool, err error) {
		node.Format([]byte)
	},	stmt)
	fmt.Printf("stmt = %+v\n", stmt)
}
