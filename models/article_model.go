package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"hhp1614/myblog/utils"
	"log"
	"strconv"
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	CreateTime int64
	//Status int // Status=0为正常，1为删除，2为冻结
}

// ---------- 添加文章 ----------
func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	SetArticleRowsNum()
	return i, err
}

// ---------- 插入一篇文章 ----------
func insertArticle(article Article) (int64, error) {
	return utils.ModifyDB("INSERT INTO article(title, tags, short, content, author, create_time) values(?,?,?,?,?,?);",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.CreateTime)
}

// ---------- 查询文章 ----------
// 根据页码查询文章
func FindArticleWithPage(page int) ([]Article, error) {
	// 从配置文件中获取每页的文章数量
	num, _ := beego.AppConfig.Int("articleListPageNum")
	page--
	fmt.Println("page ---------->", page)
	return QueryArticleWithPage(page, num)
}

// 根据id查询文章
func QueryArticleWithId(id int) Article {
	row := utils.QueryRowDB("SELECT id,title,tags,short,content,author,create_time FROM article WHERE id=" + strconv.Itoa(id))
	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createTime int64 = 0
	row.Scan(&id, &title, &tags, &short, &content, &author, &createTime)
	art := Article{id, title, tags, short, content, author, createTime}
	return art
}

/**
分页查询数据库
limit分页查询语句，
	语法：limit m，n

	m代表从多少位开始获取，与id值无关
	n代表获取多少条数据

注意limit前面没有where
*/
func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("LIMIT %d,%d", page*num, num)
	return QueryArticlesWithCon(sql)
}

func QueryArticlesWithCon(sql string) ([]Article, error) {
	sql = "SELECT id,title,tags,short,content,author,create_time FROM article " + sql
	rows, err := utils.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var artList []Article
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createTime int64 = 0

		rows.Scan(&id, &title, &tags, &short, &content, &author, &createTime)
		art := Article{id, title, tags, short, content, author, createTime}

		artList = append(artList, art)
	}
	return artList, nil
}

// ---------- 翻页 ----------
// 存储表的行数，只有自己可以更改，当文章新增或者删除时需要更新这个值
var articleRowsNum = 0

// 只有首次获取行数的时候采取统计表里的行数
func GetArticleRowsNum() int {
	if articleRowsNum == 0 {
		articleRowsNum = QueryArticleRowNum()
	}
	return articleRowsNum
}

// 查询文章的总条数
func QueryArticleRowNum() int {
	row := utils.QueryRowDB("SELECT COUNT(id) FROM article")
	num := 0
	row.Scan(&num)
	return num
}

// 设置页数
func SetArticleRowsNum() {
	articleRowsNum = QueryArticleRowNum()
}

// ---------- 按照标签查询 ----------
/*
通过标签查询首页的数据
有四种情况
	1.左右两边有&符和其他符号
	2.左边有&符号和其他符号，同时右边没有任何符号
	3.右边有&符号和其他符号，同时左边没有任何符号
	4.左右两边都没有符号
通过%去匹配任意多个字符，至少是一个
*/
func QueryArticlesWithTag(tag string) ([]Article, error) {
	sql := " WHERE tags LIKE '%&" + tag + "&%'"
	sql += " OR tags LIKE '%&" + tag + "'"
	sql += " OR tags LIKE '" + tag + "&%'"
	sql += " OR tags LIKE '" + tag + "'"
	fmt.Println(sql)
	return QueryArticlesWithCon(sql)
}

// ---------- 修改文章 ----------
func UpdateArticle(article Article) (int64, error) {
	// 数据库操作
	return utils.ModifyDB("UPDATE article SET title=?,tags=?,short=?,content=? WHERE id=?",
		article.Title, article.Tags, article.Short, article.Content, article.Id)
}

// ---------- 删除文章 ----------
func DeleteArticle(artID int) (int64, error) {
	i, err := deleteArticleWithId(artID)
	SetArticleRowsNum()
	return i, err
}

func deleteArticleWithId(artID int) (int64, error) {
	return utils.ModifyDB("DELETE FROM article WHERE id=?", artID)
}

// 查询标签，返回一个字段的列表
func QueryArticleWithParam(param string) []string {
	rows, err := utils.QueryDB(fmt.Sprintf("SELECT %s FROM article", param))
	if err != nil {
		log.Println(err)
	}
	var paramList []string
	for rows.Next() {
		arg := ""
		rows.Scan(&arg)
		paramList = append(paramList, arg)
	}
	return paramList
}
