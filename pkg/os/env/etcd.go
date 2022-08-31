package env

import (
	"os"
	"strings"
)

func GetEtcdEndpoints() []string {
	etcdEndpoints := os.Getenv("etcdEndpoints")
	return strings.Split(etcdEndpoints, ",")
}
