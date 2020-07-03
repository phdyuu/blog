package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Moment struct {
	Id int `orm:"pk;auto"`
	Title string `orm:"size(20)"`
	Content string `orm:"size(500)"`
	Class string `orm:"size(20)"`
	Img string `orm:"size(50);null"`
	Time time.Time `orm:"type:(datatime);auto_now_add"`
	Count int `orm:"default(0)"`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:gotoburnout@tcp(192.168.153.133:3306)/phdyu?charset=utf8")
	orm.RegisterModel(new(Moment))
	orm.RunSyncdb("default", false, true)


}

