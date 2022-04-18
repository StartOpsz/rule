package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func BasicSign(version string, timeStamp string, signatureNonce string, accessKeyId string,
	accessSecretId string) (string, error) {

	if version == "" || timeStamp == "" || signatureNonce == "" || accessKeyId == "" || accessSecretId == "" {
		return "", errors.New("参数异常")
	}

	// check timeStamp
	timeStampInt, err := strconv.Atoi(timeStamp)
	if err != nil {
		return "", err
	}

	timeStampTime := time.Unix(int64(timeStampInt), 0)
	interval := time.Until(timeStampTime).Seconds()
	if interval <= -120 {
		return "", errors.New("")
	}

	// compute sign
	hmacMessage := version + timeStamp + signatureNonce + accessKeyId

	mac := hmac.New(sha256.New, []byte(accessSecretId))
	mac.Write([]byte(hmacMessage))
	expectedMAC := hex.EncodeToString(mac.Sum(nil))

	return expectedMAC, nil
}

func GinSign(serviceAccessKeyId, serviceSecretKeyId string) gin.HandlerFunc {
	return func(c *gin.Context) {
		version := c.Query("Version")
		timeStamp := c.Query("TimeStamp")
		signatureNonce := c.Query("SignatureNonce")
		accessKeyId := c.Query("AccessKeyId")
		signatureToken := c.Query("SignatureToken")

		accessSecretId := serviceSecretKeyId

		// accessKeyId 对比
		if accessKeyId != serviceAccessKeyId {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"errCode": http.StatusUnauthorized, "message": "accessKeyId 参数异常."})
			return
		}

		// check version
		if version != "v1" {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"errCode": http.StatusUnauthorized, "message": "version 参数异常."})
			return
		}

		// check timeStamp
		timeStampInt, err := strconv.Atoi(timeStamp)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"errCode": http.StatusUnauthorized, "message": "TimeStamp 要求时间戳整数格式."})
			return
		}

		timeStampTime := time.Unix(int64(timeStampInt), 0)
		interval := time.Until(timeStampTime).Seconds()
		if interval <= -120 {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"errCode": http.StatusUnauthorized, "message": "TimeStamp 参数已过期."})
			return
		}

		if signatureNonce == "" {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"errCode": http.StatusUnauthorized, "message": "signatureNonce 鉴权参数异常."})
			return
		}
		// compute sign
		hmacMessage := version + timeStamp + signatureNonce + accessKeyId

		if !validMAC([]byte(hmacMessage), signatureToken, []byte(accessSecretId)) {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"errCode": http.StatusUnauthorized, "message": "鉴权失败."})
			return
		}
		c.Next()
	}
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
