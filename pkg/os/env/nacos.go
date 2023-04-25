package env

import (
	"os"
	"strings"
)

/*
nacos NameSpaceId, 未配置则为空
*/

func GetNacosNameSpaceId() string {
	return os.Getenv("nacosNameSpaceId")
}

func GetNacosEndpoints() []string {
	nacosEndpoints := os.Getenv("nacosEndpoints")
	if nacosEndpoints == "" {
		return []string{}
	}
	return strings.Split(nacosEndpoints, ",")
}
