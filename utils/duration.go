package utils

import (
	"strconv"
	"strings"
	"time"
)

var TimeTemplates = "2006-01-02 15:04:05" //常规类型

func TimeStringToGoTime(tm, format string) time.Time {
	t, err := time.ParseInLocation(format, tm, time.Local)
	if err == nil && !t.IsZero() {
		return t
	}
	return time.Time{}
}

// ParseDuration 解析时间
func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour)
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}
