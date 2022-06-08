package http

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/stretchr/testify/assert"
	"testing"
	
	nHttp "net/http"
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
		Method: Post,
		Url: "https://robot.startops.com.cn/v1",
		Body: b,
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


func TestPostReq2(t *testing.T) {
	u := User{
		Username: "qx",
		Password: "123",
	}
	
	b, err := json.Marshal(&u)
	if err != nil {
		assert.Error(t, err, "json失败")
		return
	}
	
	_, err = nHttp.Post("https://robot.startops.com.cn/v2", "application/json", bytes.NewBuffer(b))
	if err != nil {
		assert.Error(t, err, "req失败")
		return
	}
	
	
	return
}