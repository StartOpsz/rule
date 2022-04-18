package typeconvert

import "strconv"

func Str2int(s string, defaultInt int) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		return defaultInt
	}
	return res
}
