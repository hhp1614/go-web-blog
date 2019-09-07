package models

import (
	"fmt"
	"hhp1614/myblog/utils"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int // 0 正常状态，1 删除
	CreateTime int64
}

// 插入
func InsertUser(user User) (int64, error) {
	return utils.ModifyDB("INSERT INTO users(username, password, status, create_time) values (?,?,?,?);",
		user.Username, user.Password, user.Status, user.CreateTime)
}

// 按条件查询
func QueryUserWithCon(con string) int {
	sql := fmt.Sprintf("SELECT id FROM users %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

// 根据用户名查询id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("WHERE username='%s'", username)
	return QueryUserWithCon(sql)
}

// 根据用户名和密码，查询id
func QueryUserWithParam(username, password string) int {
	sql := fmt.Sprintf("WHERE username='%s' AND password='%s'", username, password)
	return QueryUserWithCon(sql)
}
