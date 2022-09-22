package constant

import "errors"

type ErrInfo struct {
	ErrCode int32
	ErrMsg  string
}

var (
	ErrDB = ErrInfo{ErrCode: 802, ErrMsg: DBMsg.Error()}
)

var (
	DBMsg = errors.New("db failed")
)
