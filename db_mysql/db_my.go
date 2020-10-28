package db_mysql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql")

var Db *sql.DB

func ConDB() {
	config := beego.AppConfig
	dbDriveername := config.String("db_driveerName") //mysql
	dbUser := config.String("db_user")               //root
	dbPassword := config.String("db_password")       //123456
	dbIp := config.String("db_ip")                   //（127.0.0.1:3306）
	dbName := config.String("db_name")               //数据库名称

	connUrl := dbUser + ":" + dbPassword + "@tcp" + dbIp + "/" + dbName + "?charset=utf8"
	db, err1 := sql.Open(dbDriveername, connUrl)
	if err1 != nil {
		//fmt.Println(err1.Error())
		panic("数据库连接错误，请检查配置")
	}
	Db = db
	fmt.Println("连接成功咯！！")
}
