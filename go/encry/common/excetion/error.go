package excetion

import (
	"encry/common/logs"
)

func ErrorRecover(msg string) {
	if r := recover(); r != nil {
		logs.Error(msg)
	}
}
