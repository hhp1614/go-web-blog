package controllers

import (
	"fmt"
	"hhp1614/myblog/models"
	"time"
)

type AddArticleController struct {
	BaseController
}

/*
当访问/add路径的时候回触发AddArticleController的Get方法
响应的页面是通过TpName
*/
func (c *AddArticleController) Get() {
	c.TplName = "write_article.html"
}

// 通过this.ServerJSON()这个方法去返回json字符串
func (c *AddArticleController) Post() {
	// 获取浏览器传输的数据，通过表单的name属性获取值
	title := c.GetString("title")
	tags := c.GetString("tags")
	short := c.GetString("short")
	content := c.GetString("content")
	fmt.Printf("title: %s, tags: %s\n", title, tags)

	// 实例化model，将它出入到数据库中
	article := models.Article{
		Id:         0,
		Title:      title,
		Tags:       tags,
		Short:      short,
		Content:    content,
		Author:     "hhp1614",
		CreateTime: time.Now().Unix(),
	}
	_, err := models.AddArticle(article)

	// 返回数据给浏览器
	var response map[string]interface{}
	if err == nil {
		// 无误
		response = map[string]interface{}{"code": 1, "message": "添加成功"}
	} else {
		response = map[string]interface{}{"code": 0, "message": "添加失败"}
	}

	c.Data["json"] = response
	c.ServeJSON()
}
