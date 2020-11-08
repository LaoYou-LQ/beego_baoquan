package models

import (
	"DataCertProject/db_mysql"
	"DataCertProject/util"
)

type User struct {
	Id       int    `form:"id"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
	Name     string `form:"name"`
	Card     string `form:"card"`
	Sex      string `form:"sex"`
}

func (u User) SeveUser() (int64, error) {
	//密码拖敏处理
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

func (u User) Querys() (*User, error) {
	util.Md5Hash(u.Password)
	row := db_mysql.Db.QueryRow("select phone , namess , card  ,sex from baoquan where phone =? and password=? ", u.Phone, u.Password)

	err := row.Scan(&u.Phone, &u.Name, &u.Card, &u.Sex)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func QueryUserByPhone(phone string) (*User, error) {
	row := db_mysql.Db.QueryRow("select phone,namess ,card, sex from baoquan where phone=?", phone)
	var user User
	err := row.Scan(&user.Phone, &user.Name, &user.Card, &user.Sex)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (u User) Update() (int64, error) {
	rs, err := db_mysql.Db.Exec("update baoquan set namess=?, card=?, sex=? where phone=?", u.Name, u.Card, u.Sex, &u.Phone)
	if err != nil {
		return -1, err
	}
	return rs.RowsAffected()
}
