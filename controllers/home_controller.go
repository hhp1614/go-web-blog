package controllers

import (
	"fmt"
	"hhp1614/myblog/models"
)

type HomeController struct {
	BaseController
}

/**
 * 请求：http://localhost:8080/
 * 请求类型：Get
 * 请求描述：
 */
func (c *HomeController) Get() {
	// 分页
	// http://localhost:8089
	// 标签
	// http://localhost:8089?tag=web

	tag := c.GetString("tag")
	fmt.Println("tag:", tag)
	page, _ := c.GetInt("page")
	var artList []models.Article

	if len(tag) > 0 {
		// 按照指定的标签搜索
		artList, _ = models.QueryArticlesWithTag(tag)
		c.Data["HasFooter"] = false
	} else {
		if page <= 0 {
			page = 1
		}
		// 设置分页
		artList, _ = models.FindArticleWithPage(page)
		c.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
		c.Data["HasFooter"] = true
	}

	fmt.Println("IsLogin:", c.IsLogin, c.LoginUser)
	c.Data["Content"] = models.MakeHomeBlocks(artList, c.IsLogin)

	c.TplName = "home.html"
}
