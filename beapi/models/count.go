package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id        int       `orm: "pk;auto"`
	Name      string    `orm: "size(20)"`
	Update time.Time `orm: "column(Update);null;auto_now;type(datetime)"`
}

func init() {
	// 注册表
	orm.RegisterModel(new(User))
	orm.DefaultTimeLoc = time.UTC
}

func (u *User) TableEngine() string {
	return "INNODB"
}

func (u *User) TableName() string {
	return "UserTest"
}

func AddObject(ob *User) int64 {
	var o = orm.NewOrm()
	//beego.Debug(ob)
	ob.Update = time.Now()
	i, e := o.Insert(ob)
	if e != nil {
		beego.Error(e.Error())
	}
	return i
}

func IntoTable(name string) int64 {
	o := orm.NewOrm()
	if err := o.Using("default");err != nil {
		beego.Error(err.Error())
	}
	user := new(User)
	user.Name = name
	user.Update = time.Now()
	i, _ := o.Insert(user)
	//if err := o.Commit(); err != nil {
	//	beego.Error(err.Error())
	//}
	return i
}

//func RetOne(id int) string {
//	user := new(User)
//	err := o.QueryTable(user).Filter("id",id).One(&user)
//	if err == orm.ErrMultiRows {
//		panic("return multi Rows")
//	}
//	if err == orm.ErrNoRows {
//		panic("Not row")
//	}
//	return user.Name
//}
//
//func RetAll(id int) ([]User,int64) {
//	var users []User
//	user := new(User)
//	num, err := o.QueryTable(user).Filter("id",id).All(&users)
//	if err == orm.ErrNoRows{
//		panic("Not rows")
//	}
//	return users,num
//}
