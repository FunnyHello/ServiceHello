package main

import (
	"./controller"
	"./hzx"
	"log"
	"net/http"
	"time"
)

func main() {
	initUtils()

	//初始化数据库
	hzx.InitDB()
	//创建表
	hzx.CreateTable()
	//定义http服务器的参数
	server := &http.Server{
		Addr:        ":8080",
		Handler:     hzx.Router,
		ReadTimeout: 5 * time.Second,
	}
	RegiterRouter(hzx.Router)
	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

//注册路由
func RegiterRouter(handler *hzx.RouterHandler) {
	//初始化用户控制器
	new(controller.UserConterller).Router(handler)
}

func initUtils() {
	go func() {
		utils.Twinkle()
	}()

	//以并发的方式调用匿名函数func
	//go func() {
	//	for {
	//		utils.GetTemperature()
	//		//延时10秒读取一次
	//		time.Sleep(10 * time.Second)
	//	}
	//
	//}()

}
