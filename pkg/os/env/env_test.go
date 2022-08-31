package env

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetNacosEndpoints(t *testing.T) {
	err := os.Setenv("nacosEndpoints", "127.0.0.1:8848,127.0.0.2:8848,127.0.0.3:8848")
	
	assert.Nil(t, err)
	
	nacosEndpoints := GetNacosEndpoints()
	
	fmt.Println("nacosEndpoints: ", nacosEndpoints)
}

func TestGetOpenTelemetryTraceEnable(t *testing.T) {
	trace := GetOpenTelemetryTraceEnable()
	fmt.Println("trace: ", trace)
	
	err := os.Setenv("openTelemetryTraceEnable", "true")
	assert.Nil(t, err)
	
	trace = GetOpenTelemetryTraceEnable()
	fmt.Println("trace: ", trace)
}
