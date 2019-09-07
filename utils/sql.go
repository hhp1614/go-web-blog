package utils

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func InitMysql() {
	fmt.Println("InitMysql......")
	driverName := beego.AppConfig.String("driverName")

	// 注册数据库驱动
	//orm.RegisterDriver(driverName, orm.DRMySQL)

	// 数据库连接
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	//dbConn := "admin:123456@tcp(192.168.0.43:3306)/my_blog_web?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	database, err := sql.Open(driverName, dbConn)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		db = database
		// 创建用户表
		CreateTableWithUser()
		// 创建文章表
		CreateTableWithArticle()
	}
}

// 操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

// 创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		create_time INT(10)
		);`
	ModifyDB(sql)
}

// 创建文章表
func CreateTableWithArticle() {
	sql := `CREATE TABLE IF NOT EXISTS article(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		title VARCHAR(30),
		author VARCHAR(20),
		tags VARCHAR(30),
		short VARCHAR(255),
		content LONGTEXT,
		create_time INT(10)
		);`
	ModifyDB(sql)
}

// 查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}
