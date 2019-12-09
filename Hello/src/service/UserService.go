package service

import (
	"../dao"
	"../entity"
	"time"
)

//用户服务
type UserService struct {
}

var userDao = new(dao.UserDao)

//插入用户数据
func (p *UserService) Insert(userName, password string) int64 {
	return userDao.Insert(&entity.User{UserName: userName, Password: password, CreateTime: time.Now()})
}

//用户名查用户数据
func (p *UserService) SelectUserByName(userName string) []entity.User {
	return userDao.SelectUserByName(userName)
}

//id查用户
func (p *UserService) SelectUserByUserId(userId int64) []entity.User {
	return userDao.SelectUserByUserId(userId)
}
