package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func GetTemperatureHumidity() (float64, float64) {
	//执行一个shell命令并获取结果
	//https://blog.csdn.net/u013807095/article/details/77458795
	ommand := "child_process"
	cmd := exec.Command("/usr/src/dht22/current.py", "-c", ommand)
	bytes,err := cmd.Output()
	if err != nil {
		log.Println(err)
	}
	resp := string(bytes)
	log.Println(resp)
	//解析传感器数据
	result := strings.Split(string(resp), "\n")
	temperature, temperatureErr := strconv.ParseFloat(result[0], 64)
	if temperatureErr != nil {
		log.Println(temperatureErr)
	}
	humidity, humidityErr := strconv.ParseFloat(result[1], 64)
	if humidityErr != nil {
		log.Println(humidityErr)
	}
	return temperature, humidity
}

//获取温度//http://shumeipai.nxez.com/2013/10/03/raspberry-pi-temperature-sensor-monitors.html
func GetTemperature() float64 {
	path := "/sys/bus/w1/devices/28-00000494cb79/w1_slave"
	//使用ioutil读取文件的所有内容（内容过大不要用此方法）
	text, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	//用换行符分割字符串数组，并取第二行
	secondline := strings.Split(string(text), "\n")[1]
	//用空格分割字符串成数组，并取最后一个，即t=23000
	temperaturedata := strings.Split(secondline, " ")[9]
	//取t=后面的数值，并转换为浮点型
	temperature, temperatureErr := strconv.ParseFloat(temperaturedata[2:], 64)
	if temperatureErr != nil {
		log.Println(temperatureErr)
	}
	//转换单位为摄氏度
	temperature = temperature / 1000
	//保留3位小数
	log.Println("温度为（摄氏度）：", temperature)
	return temperature
}

//小灯闪烁
func Twinkle() {
	gpio := "/sys/class/gpio/export"         //系统的GPIO调用文件，为了暴露GPIO操作文件
	out := "/sys/class/gpio/gpio2/direction" //具体的GPIO2操作输入输出文件，可以修改gpio2来操作对应GPIO口
	value := "/sys/class/gpio/gpio2/value"   //输出为1否则为0
	input("2", gpio)
	fmt.Printf("打开2GPIO口\n")
	time.Sleep(1 * time.Second)
	input("out", out)
	fmt.Printf("更改模式为输出\n")
	time.Sleep(1 * time.Second)
	for i := 1; i <= 100; i++ {
		input("1", value)
		fmt.Printf("输出高电频开灯%d次\n", i)
		time.Sleep(1 * time.Second)
		input("0", value)
		fmt.Printf("输出低电频关灯%d次\n", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("程序结束!!\n")
}
func input(message string, files string) {
	file, err := os.OpenFile(files, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	//把字符串写入文件
	file.WriteString(message)
	file.Close()
}
