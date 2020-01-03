package main

import (
	_ "go-cms/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	logs.SetLevel(beego.LevelInformational)//设置日志等级
	//tips: 要用``  而不是单引号
	logs.SetLogger("file",`{"filename":"logs/test.log"}`)
	beego.Run()
}

