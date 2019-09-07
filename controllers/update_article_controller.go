package controllers

import (
	"fmt"
	"hhp1614/myblog/models"
)

type UpdateArticleController struct {
	BaseController
}

// 当访问/update路径的时候回触发Get()方法，响应的页面是通过TplName这个属性指定返回给客户端的页面
func (c *UpdateArticleController) Get() {
	id, _ := c.GetInt("id")
	fmt.Println("update id:", id)

	// 获取id所对应的文章信息
	art := models.QueryArticleWithId(id)
	c.Data["Title"] = art.Title
	c.Data["Tags"] = art.Tags
	c.Data["Short"] = art.Short
	c.Data["Content"] = art.Content
	c.Data["Id"] = art.Id
	c.TplName = "write_article.html"
}

func (c *UpdateArticleController) Post() {
	id, _ := c.GetInt("id")
	fmt.Println("postId", id)

	// 获取浏览器传输的数据，通过表单的name属性获取值
	title := c.GetString("title")
	tags := c.GetString("tags")
	short := c.GetString("short")
	content := c.GetString("content")

	// 实例化model，修改数据库
	art := models.Article{
		Id:         id,
		Title:      title,
		Tags:       tags,
		Short:      short,
		Content:    content,
		Author:     "",
		CreateTime: 0,
	}
	_, err := models.UpdateArticle(art)

	// 返回数据给浏览器
	if err == nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "更新成功"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "更新失败"}
	}
	c.ServeJSON()
}
