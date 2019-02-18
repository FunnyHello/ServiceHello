package controller

import (
	"../hzx"
	"../utils"
	"log"
	"net/http"
	"strconv"
	"time"
)

type RaspberryPiConterller struct {
	hzx.ApiController
}

func (p *RaspberryPiConterller) Router(router *hzx.RouterHandler) {
	router.Router("/turn_on_the_light", p.turnOnTheLight)
}
func (p *RaspberryPiConterller) turnOnTheLight(w http.ResponseWriter, r *http.Request) {
	//go func() {
	//	utils.Twinkle()
	//}()

	//go func() {
	//	for {
	//		utils.GetTemperature()
	//		//延时10秒读取一次
	//		time.Sleep(10 * time.Second)
	//	}
	//
	//}()
	go func() {
		for {
			temperature, humidity := utils.GetTemperatureHumidity()
			log.Println("温度为（摄氏度）：",temperature)
			log.Println("湿度："+ strconv.FormatFloat(humidity, 'E', -1, 64))
			time.Sleep(1 * time.Second)
		}
	}()

}
