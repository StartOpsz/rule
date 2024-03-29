package timestamp

import (
	"strings"
	"time"
)

/*
timeString: "2006-01-02 15:04:05 +0800 CST"
*/

func KubernetesStringToTime(timeString string) (time.Time, error) {
	//loc , err := time.LoadLocation("Asia/Shanghai")
	
	loc, err := time.LoadLocation("CST")
	if err != nil {
		return time.Parse("2006-01-02 15:04:05 +0800 CST", timeString)
	}
	
	return time.ParseInLocation("2006-01-02 15:04:05 +0800 CST", timeString, loc)
}

/*
timeString: "2021-06-10T06:33:23.168062118Z"
*/

func PrometheusStringToTime(timeString string) (time.Time, error) {
	timeString1 := strings.Split(timeString, ".")[0]
	loc, err := time.LoadLocation("CST")
	if err != nil {
		return time.Parse("2006-01-02T15:04:05", timeString1)
	}
	
	return time.ParseInLocation("2006-01-02T15:04:05", timeString, loc)
}

// utc string 格式: 2011-06-01T15:00:00Z

func UTCStringToTime(timeString string) (time.Time, error) {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		return time.Now(), err
	}
	
	return time.ParseInLocation("2006-01-02T15:04:05Z", timeString, loc)
}

func StringToTime(timeString string, location string) (time.Time, error) {
	loc, err := time.LoadLocation(location)
	//loc, err := time.LoadLocation("CST")
	if err != nil {
		return time.Now(), err
		//return time.Parse("2006-01-02 15:04:05", timeString)
	}
	
	return time.ParseInLocation("2006-01-02 15:04:05", timeString, loc)
}
