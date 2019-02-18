package main

import (
	"./controller"
	"./hzx"
	"./utils"
	"log"
	"net/http"
	"time"
)

func main() {

	go func() {
		for {
			temperature, humidity := utils.GetTemperatureHumidity()
			log.Println("++++++++++++++++++")
			log.Println("温度：",temperature)
			log.Println("湿度：",humidity)
			time.Sleep(1 * time.Second)
		}
	}()

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
	//初始化用户控制器
	new(controller.RaspberryPiConterller).Router(handler)
}
