package consts

import "errors"

var (
	StarPostRepeatedErr = errors.New("不允许重复点赞")
	DBErrNotFound       = errors.New("db Can't be nil")
)
