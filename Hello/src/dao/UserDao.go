package dao

import (
	"../entity"
	"../hzx"
	"log"
)

type UserDao struct {
}

func (p *UserDao) Insert(user *entity.User) int64 {
	result, err := hzx.DB.Exec("INSERT INTO user(`user_name`,`password`,`create_time`) value(?,?,?)", user.UserName, user.Password, user.CreateTime)
	if err != nil {
		log.Println(err)
		return 0
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0
	}
	return id
}

func (p *UserDao) SelectUserByName(userName string) []entity.User{
	rows,err := hzx.DB.Query("SELECT * FROM user WHERE user_name = ?",userName)
	if err != nil {
		log.Println(err)
		return nil
	}
	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.UserId,&user.UserName,&user.Password,&user.CreateTime)
		if err != nil{
			log.Println(err)
			continue
		}
		users = append(users,user)
	}
	rows.Close()
	return users
}
func (p *UserDao) SelectUserByUserId(userId int64) []entity.User{
	rows,err := hzx.DB.Query("SELECT * FROM user WHERE user_id = ?",userId)
	if err != nil {
		log.Println(err)
		return nil
	}
	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.UserId,&user.UserName,&user.Password,&user.CreateTime)
		if err != nil{
			log.Println(err)
			continue
		}
		users = append(users,user)
	}
	rows.Close()
	return users
}
