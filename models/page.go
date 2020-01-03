package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
type Page struct {
	Id int
	Website string
	Email  string
}

func init(){
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/laravel?charset=utf8")
	orm.RegisterModel(new(Page))
}

func GetPage() Page {
	o := orm.NewOrm();
	p := Page{Id:1}
	err := o.Read(&p)
	fmt.Print(err)

	return p
}

func UpdatePage(){
	p := Page{Id:1, Email:"lastemail@123.com"}
	o := orm.NewOrm()
	//如果只改一个字段,需要声明下
	o.Update(&p, "Email")
	//o.Delete()
	//o.Insert()
}