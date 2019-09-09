package controllers

import (
	"fmt"
	"hhp1614/myblog/models"
)

type TagsController struct {
	BaseController
}

func (c *TagsController) Get() {
	tags := models.QueryArticleWithParam("tags")
	fmt.Println(models.HandleTagsListData(tags))
	c.Data["Tags"] = models.HandleTagsListData(tags)
	c.TplName = "tags.html"
}
