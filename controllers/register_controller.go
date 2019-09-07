package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"hhp1614/myblog/models"
	"hhp1614/myblog/utils"
	"time"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "register.html"
}

// 处理注册
func (c *RegisterController) Post() {
	// 获取表单信息
	username := c.GetString("username")
	password := c.GetString("password")
	rePassword := c.GetString("rePassword")
	fmt.Println(username, password, rePassword)

	// 注册之前先判断该用户名是否已经被注册，如果已经注册，返回错误
	id := models.QueryUserWithUsername(username)
	fmt.Println("id:", id)
	if id > 0 {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "用户名已存在"}
		c.ServeJSON()
		return
	}

	// 注册用户名和密码
	// 村塾的密码是md5后的数据，那么在登录的验证的时候，也是需要将用户的密码md5之后和数据库里面的里面的密码进行判断
	password = utils.MD5(password)
	fmt.Println("md5后:", password)

	user := models.User{
		Id:         0,
		Username:   username,
		Password:   password,
		Status:     0,
		CreateTime: time.Now().Unix(),
	}
	_, err := models.InsertUser(user)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "注册成功"}
	}
	c.ServeJSON()
}
