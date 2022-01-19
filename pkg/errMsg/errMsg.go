package errMsg

import "fmt"

func NewMsg(requestId, msg string) string {
	return fmt.Sprintf("requestId: %s, message: %s", requestId, msg)
}
