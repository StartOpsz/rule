package sign

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/go-playground/validator"
	"strconv"
	"time"
)

// 基于 AccessKey & AccessSecretKey 对称加密

type KeySign struct {
	Version        string `validate:"required"`
	TimeStamp      int64  `validate:"required"` // 秒
	SignatureNonce string `validate:"required"`
	AccessKeyId    string `validate:"required"`
	AccessSecretId string `validate:"required"`
}

func (k KeySign) Sign() (string, error) {
	
	valid := validator.New()
	err := valid.Struct(k)
	if err != nil {
		return "", err
	}
	// timestamp check
	timeStampTime := time.Unix(k.TimeStamp, 0)
	interval := time.Until(timeStampTime).Seconds()
	if interval <= -120 {
		return "", errors.New("timestamp expired")
	}
	
	// compute sign
	timeStampString := strconv.Itoa(int(k.TimeStamp))
	hmacMessage := k.Version + timeStampString + k.SignatureNonce + k.AccessKeyId
	mac := hmac.New(sha256.New, []byte(k.AccessSecretId))
	mac.Write([]byte(hmacMessage))
	expectedMAC := hex.EncodeToString(mac.Sum(nil))
	
	return expectedMAC, nil
}

func validMAC(message []byte, messageMAC string, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := hex.EncodeToString(mac.Sum(nil))
	if expectedMAC == messageMAC {
		return true
	}
	return false
}
