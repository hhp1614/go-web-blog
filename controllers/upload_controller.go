package controllers

import (
	"fmt"
	"hhp1614/myblog/models"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type UploadController struct {
	BaseController
}

func (c *UploadController) Post() {
	fmt.Println("file upload...")
	fileData, fileHeader, err := c.GetFile("upload")
	if err != nil {
		c.responseErr(err)
		return
	}
	fmt.Println("name:", fileHeader.Filename, fileHeader.Size)
	now := time.Now()
	fmt.Println("ext:", filepath.Ext(fileHeader.Filename))
	fileType := "other"
	// 判断后缀为图片的文件，如果时图片才存入到数据库中
	fileExt := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" {
		fileType = "img"
	}
	// 文件夹路径
	fileDir := fmt.Sprintf("static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	// ModePerm时0777，这样拥有该文件夹路径的执行权限
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		c.responseErr(err)
		return
	}
	// 文件路径
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fileHeader.Filename)
	filePathStr := filepath.Join(fileDir, fileName)
	desFile, err := os.Create(filePathStr)
	if err != nil {
		c.responseErr(err)
		return
	}
	// 将浏览器客户端上传的文件拷贝到本地路径的文件里面
	_, err = io.Copy(desFile, fileData)
	if err != nil {
		c.responseErr(err)
		return
	}
	if fileType == "img" {
		album := models.Album{
			Id:         0,
			Filepath:   filePathStr,
			Filename:   fileName,
			Status:     0,
			CreateTime: timeStamp,
		}
		models.InsertAlbum(album)
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "message": "上传成功"}
	c.ServeJSON()
}

func (c *UploadController) responseErr(err error) {
	c.Data["json"] = map[string]interface{}{"code": 0, "message": err}
	c.ServeJSON()
}
