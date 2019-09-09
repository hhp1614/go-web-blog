package controllers

import (
	"hhp1614/myblog/models"
	"log"
)

type AlbumController struct {
	BaseController
}

func (c *AlbumController) Get() {
	albums, err := models.FindAllAlbums()
	if err != nil {
		log.Println(err)
	}
	c.Data["Album"] = albums
	c.TplName = "album.html"
}
