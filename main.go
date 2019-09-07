package main

import (
	"github.com/astaxie/beego"
	_ "hhp1614/myblog/routers"
	"hhp1614/myblog/utils"
)

func main() {
	utils.InitMysql()
	beego.Run()
}
