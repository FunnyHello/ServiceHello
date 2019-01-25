package controller

import (
	"../constant"
	"../entity"
	"../hzx"
	"../service"
	"../utils"
	"encoding/json"
	"log"
	"net/http"
)

type UserConterller struct {
	hzx.ApiController
}

var userService = new(service.UserService)

func (p *UserConterller) Router(router *hzx.RouterHandler) {
	router.Router("/register", p.register)
	router.Router("/login", p.login)
}

//注册
func (p *UserConterller) register(w http.ResponseWriter, r *http.Request) {
	resultMold := &entity.ResultMold{}
	////解析Header
	log.Println("Header参数：")
	log.Println(r.Header)
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			log.Println("键：" + k + "  值：" + v[0])
		}
	}
	//解析Header
	log.Println("RequestURI参数：")
	log.Println(r.RequestURI)
	//获取http传入的data
	data := r.PostFormValue("data")
	//json转实体类
	user := &entity.User{}
	err := json.Unmarshal([]byte(data), user)
	if err != nil {
		log.Panic(err)
		resultMold.Code = 200
		resultMold.Message = "提交数据错误"
		resultMold.Data = utils.EntityToJson(entity.RegisterEntity{Status: 4})
		hzx.ResultOk(w, utils.EntityToJson(resultMold))
		return
	}
	log.Println(*p)
	userName := user.UserName
	password := user.Password

	if utils.IsEmpty(userName) || utils.IsEmpty(password) {

		resultMold.Code = 200
		resultMold.Message = "用户名或密码不能为空"
		resultMold.Data = utils.EntityToJson(entity.RegisterEntity{Status: 3})
		hzx.ResultOk(w, utils.EntityToJson(resultMold))
		return
	}
	//根据用户名查用户信息
	dbUsers := userService.SelectUserByName(userName)
	//查到数据了
	if len(dbUsers) != 0 {
		if dbUsers[0].UserName == userName {
			resultMold.Code = 200
			resultMold.Message = "此账号已注册"
			resultMold.Data = utils.EntityToJson(entity.RegisterEntity{Status: 5})
			hzx.ResultOk(w, utils.EntityToJson(resultMold))
			return
		}
	}

	//把数据插入数据库
	id := userService.Insert(userName, password)
	if id <= 0 {
		resultMold.Code = 200
		resultMold.Message = "注册失败"
		resultMold.Data = utils.EntityToJson(entity.RegisterEntity{Status: 1})
		hzx.ResultOk(w, utils.EntityToJson(resultMold))
		return
	}
	resultMold.Code = 200
	resultMold.Message = "注册成功"
	resultMold.Data = utils.EntityToJson(entity.RegisterEntity{Status: 0})
	hzx.ResultOk(w, utils.EntityToJson(resultMold))
}

//登陆
func (p *UserConterller) login(w http.ResponseWriter, r *http.Request) {
	resultMold := &entity.ResultMold{}


	data := r.PostFormValue("data")
	user := &entity.User{}
	err := json.Unmarshal([]byte(data), user)
	if err != nil {
		log.Panic(err)
		resultMold.Code = 200
		resultMold.Message = "提交数据错误"
		resultMold.Data = utils.EntityToJson(entity.RegisterEntity{Status: 4})
		hzx.ResultOk(w, utils.EntityToJson(resultMold))
		return
	}
	userName := user.UserName
	password := user.Password
	if utils.IsEmpty(userName) || utils.IsEmpty(password) {
		resultMold.Code = 200
		resultMold.Message = "用户名或密码不能为空"
		resultMold.Data = utils.EntityToJson(entity.RegisterEntity{Status: 3})
		hzx.ResultOk(w, "用户名或密码不能为空")
		return
	}
	users := userService.SelectUserByName(userName)
	if len(users) == 0 {
		resultMold.Code = 200
		resultMold.Message = "用户不存在"
		resultMold.Data = utils.EntityToJson(entity.RegisterEntity{Status: 6})
		hzx.ResultOk(w, "用户不存在")
		return
	}
	if users[0].Password != password {
		resultMold.Code = 200
		resultMold.Message = "密码错误"
		resultMold.Data = utils.EntityToJson(entity.RegisterEntity{Status: 7})
		hzx.ResultOk(w, "密码错误")
		return
	}

	//session
	session := hzx.GlobalSession().SessionStart(w, r)
	session.Set(constant.KEY_USER, &users[0])

	resultMold.Code = 200
	resultMold.Message = "登录成功"
	resultMold.Data = utils.EntityToJson(entity.RegisterEntity{Status: 0})
	hzx.ResultOk(w, "登录成功")
}
