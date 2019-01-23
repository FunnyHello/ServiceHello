package hzx

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

//数据库的配置
const (
	username   = "root"
	password   = "123456"
	ip         = "127.0.0.1"
	port       = "3306"
	dbName     = "hello"
	driverName = "mysql"
)

//DB数据库连接池
var DB *sql.DB

//初始化数据库
func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=uft8"
	//注意：要想解析time.Time类型，必须要设置parseTime=True
	path := strings.Join([]string{username, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8&parseTime=True&loc=Local"}, "")
	//打开数据库，前者是驱动名，所以要导入:_"github.com/go-sql-driver/mysql"
	DB, _ = sql.Open(driverName, path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		log.Panic(err)
	}
	log.Println("数据库连接成功")
}

//创建表
func CreateTable() {
	userTable := "CREATE TABLE IF NOT EXISTS `user`(" +
		"`user_id` INT UNSIGNED AUTO_INCREMENT ," +
		"`user_name` VARCHAR(20) NOT NULL," +
		"`password` VARCHAR(40) NOT NULL," +
		"`create_time` DATETIME," +
		"PRIMARY KEY ( `user_id` )" +
		")ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=100000;"

	//执行数据库语句
	_, err := DB.Exec(userTable)
	if err != nil {
		log.Panic(err)
	}
}
