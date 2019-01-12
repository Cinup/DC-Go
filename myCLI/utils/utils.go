package utils

import (
	"strconv"
	"strings"
)

//去掉返回值HTTP的头部
func GetBody(result []byte) (body []byte) {
	for i := 0; i <= len(result)-4; i++ {
		if result[i] == 91 && result[i+1] == 123 {
			body = result[i:]
			break
		}
	}
	return
}

//计算镜像的大小和单位
func CalSize(size int64) (result string) {
	unit := 0
	for ; size > 1024; size /= 1000 {
		unit++
	}
	result = strconv.Itoa(int(size))
	switch unit {
	case 0:
		result = result + "Byte"
	case 1:
		result = result + "KB"
	case 2:
		result = result + "MB"
	case 3:
		result = result + "GB"
	}
	return
}

//
func GetTag(s string) (tag string) {
	return s[strings.Index(s, ":")+1:]
}
