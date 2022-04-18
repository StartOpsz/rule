package timestamp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKubernetesStringToTime(t *testing.T) {
	r, err := KubernetesStringToTime("2021-01-01 10:00:00 +0800 CST")

	assert.Equal(t, "2021-01-01 10:00:00", r.Format("2006-01-02 15:04:05"),
		"转换 Kubernetes 时间结果异常")

	assert.Nil(t, err, "转换 Kubernetes 时间字符串为 Time 失败")
}

func TestPrometheusStringToTime(t *testing.T) {
	r, err := PrometheusStringToTime("2021-01-01T10:00:00.168062118Z")

	assert.Nil(t, err, "转换 Prometheus 时间字符串为 Time 失败")
	assert.Equal(t, "2021-01-01 10:00:00", r.Format("2006-01-02 15:04:05"),
		"转换 Prometheus 时间结果异常")
}
