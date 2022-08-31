package env

import (
	"os"
	"strings"
)

func GetOpenTelemetryTraceEnable() bool {
	openTelemetryTraceEnable := strings.ToLower(os.Getenv("openTelemetryTraceEnable"))
	
	switch openTelemetryTraceEnable {
	case "true":
		return true
	case "enable":
		return true
	case "on":
		return true
	case "yes":
		return true
	default:
		return false
	}
}

/*
获取 OpenTelemetry 终端地址
*/

func GetOpenTelemetryEndpoint() string {
	return os.Getenv("openTelemetryEndpoint")
}
