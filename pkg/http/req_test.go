package http

import (
	"encoding/json"
	"fmt"
	
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetReq(t *testing.T) {
	req := Req{
		Method: Get,
		Url:    "https://api.dingtalk.com/v1.0/oauth2/userAccessToken",
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
		Url:    "https://api.dingtalk.com/v1.0/oauth2/userAccessToken",
	}
	
	rb, err := req.Do()
	if err != nil {
		assert.Error(t, err, "请求失败")
		return
	}
	fmt.Println("resp Body: ", string(rb.Body))
	fmt.Println("resp StatusCode: ", rb.StatusCode)
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestPostReq1(t *testing.T) {
	u := User{
		Username: "qx",
		Password: "123",
	}
	
	b, err := json.Marshal(&u)
	
	if err != nil {
		assert.Error(t, err, "json失败")
		return
	}
	
	header := make(map[string]string)
	header["contentType"] = "application/json"
	header["user-agent"] = "startops"
	
	req := Req{
		Method:  Post,
		Url:     "https://robot.startops.com.cn/v1",
		Body:    b,
		Headers: header,
	}
	
	rb, err := req.Do()
	if err != nil {
		assert.Error(t, err, "请求失败")
		return
	}
	fmt.Println("resp Body: ", string(rb.Body))
	fmt.Println("resp StatusCode: ", rb.StatusCode)
}

func TestGetHttpCertificate(t *testing.T) {
	GetHttpCertificate("https://139.9.122.48:5443")
}

func TestGetHttpCertificate1(t *testing.T) {
	GetHttpCertificate("https://api.startops.com.cn")
}
