package entity

import "time"

type User struct {
	UserId     uint    `json:"userId"`
	UserName   string    `json:"userName"`
	Password   string    `json:"password"`
	CreateTime time.Time `json:"createTime"`
}
