package models

import "hhp1614/myblog/utils"

type Album struct {
	Id         int
	Filepath   string
	Filename   string
	Status     int
	CreateTime int64
}

// 插入图片
func InsertAlbum(album Album) (int64, error) {
	return utils.ModifyDB("INSERT INTO album(filepath,filename,status,create_time) VALUES(?,?,?,?);",
		album.Filepath, album.Filename, album.Status, album.CreateTime)
}

// 查询图片
func FindAllAlbums() ([]Album, error) {
	rows, err := utils.QueryDB("SELECT id,filepath,filename,status,create_time FROM album")
	if err != nil {
		return nil, err
	}
	var albums []Album
	for rows.Next() {
		id := 0
		filepath := ""
		filename := ""
		status := 0
		var createTime int64 = 0
		rows.Scan(&id, &filepath, &filename, &status, &createTime)
		album := Album{id, filepath, filename, status, createTime}
		albums = append(albums, album)
	}
	return albums, nil
}
