package models

import (
	"DataCertProject/db_mysql"
	"DataCertProject/util"
	"fmt"
)

type User struct {
	Id       int    `form:"id"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

func (u User) SeveUser() (int64, error) {
	fmt.Println("执行了")
	//密码拖密处理
	 //获得结构体user中的用户密码并粉碎
	util.Md5Hash(u.Password)
	//执行数据库操作
	row, err := db_mysql.Db.Exec("insert into baoquan(phone,password)"+" values(?,?)", u.Phone, u.Password)
	if err != nil {
		return -1, err
	}
	//返回受影响的行数
	num, err := row.RowsAffected()
	if err != nil {
		return -1, err
	}
	return num, nil
}

func (u User) Querys() (*User ,error) {
	util.Md5Hash(u.Password)
	row :=db_mysql.Db.QueryRow("select phone from baoquan where phone =? and password=? ", u.Phone, u.Password)

	err:=row.Scan(&u.Phone)
	if err !=nil {
	return nil ,err
}
	return &u, nil
}
