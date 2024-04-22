package conf

import (
	"fisco-drug-trace/model"
	"fisco-drug-trace/sdk"
	"fisco-drug-trace/util"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	// 初始化合约
	sdk.InitContract()
}
