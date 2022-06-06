package task

import "time"

type ExecTime struct {
	StartTime string
	EndTime   string
	CostTime  int64
}


func ExecTimeResult(startTime time.Time) ExecTime {
	endTime := time.Now()
	return ExecTime{
		StartTime: startTime.String(),
		EndTime: endTime.String(),
		CostTime: endTime.Sub(startTime).Milliseconds(),
	}
}

