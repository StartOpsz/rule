package http

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetReq(t *testing.T) {
	req := Req{
		Method: Get,
		Url: "https://api.dingtalk.com/v1.0/oauth2/userAccessToken",
	}
	
	rb, err := req.Do()
	if err != nil {
		assert.Error(t, err, "请求失败")
		return
	}
	fmt.Println("resp Body: ", string(rb.Body))
	fmt.Println("resp StatusCode: ", rb.StatusCode)
}



func TestPostReq(t *testing.T) {
	req := Req{
		Method: Post,
		Url: "https://api.dingtalk.com/v1.0/oauth2/userAccessToken",
	}
	
	rb, err := req.Do()
	if err != nil {
		assert.Error(t, err, "请求失败")
		return
	}
	fmt.Println("resp Body: ", string(rb.Body))
	fmt.Println("resp StatusCode: ", rb.StatusCode)
}
