package helper

import "strconv"

// Int64InSlice is used for check if value of int64 contains in array of int64
func Int64InSlice(val int64, arr []int64) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}

// Int64SliceToStringSlice converts []int64 to []string
func Int64SliceToStringSlice(arr []int64) []string {
	var txt string
	var data []string
	for _, v := range arr {
		txt = strconv.FormatInt(v, 10)
		data = append(data, txt)
	}
	return data
}
