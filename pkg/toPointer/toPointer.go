package toPointer

func StrPoint(v string) *string {
	return &v
}

func Int32Point(v int32) *int32 {
	return &v
}

func Int64Point(v int64) *int64 {
	return &v
}
