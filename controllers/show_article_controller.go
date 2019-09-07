package controllers

import (
	"fmt"
	"hhp1614/myblog/models"
	"hhp1614/myblog/utils"
	"strconv"
)

type ShowArticleController struct {
	BaseController
}

func (c *ShowArticleController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	fmt.Println("id:", id)
	// 获取id对应的文章的信息
	art := models.QueryArticleWithId(id)
	c.Data["Title"] = art.Title
	//c.Data["Content"] = art.Content
	c.Data["Content"] = utils.SwitchMarkdownToHtml(art.Content)
	c.TplName = "show_article.html"
}
