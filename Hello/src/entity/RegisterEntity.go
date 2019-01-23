package entity

type RegisterEntity struct {
	//0 成功 1 失败 2 重复 3 参数为空 4 提交数据格式错误 5 用户已注册 6 用户不存在 7 密码错误
	Status int64
}
