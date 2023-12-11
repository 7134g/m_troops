package xerr

const (
	SvcCodeParams    uint32 = 4000
	SvcCodeForbidden uint32 = 4003

	SvcCodeError     uint32 = 1001 // 默认不知名错误
	SvcCodeDbError   uint32 = 1002 // 数据库错误
	SvcCodeAuthError uint32 = 1003 // 验证错误
)

var ErrNotFound = &CodeError{errCode: 10000, errMsg: "无数据"}
