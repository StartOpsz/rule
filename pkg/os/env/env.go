package env

import (
	"os"
)

/*
获取程序运行环境，一般配置为: dev, beta, gray, prod
*/

func GetEnvironment() string {
	environment := os.Getenv("environment")
	if environment == "" {
		environment = "default"
	}
	return environment
}

/*
获取程序运行区域，一般配置为服务器所在的地区信息
*/

func GetRegion() string {
	region := os.Getenv("region")
	if region == "" {
		region = "default"
	}
	return region
}

func GetConfigPath() string {
	return os.Getenv("configPath")
}
