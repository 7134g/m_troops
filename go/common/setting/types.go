package setting

/*
新增配置类型只允许下列格式，否则反射失败
redis tag 没填则自动使用json

	type allSetting struct {
		name struct{
			newConfig `json:"new_config" redis:"redis hash结构下的field" default:"默认值" description:"新配置" category:"分组名"`
		}
	}
*/
type allSetting struct {
	// 用于测试类型
	testField
}

// testField 用于测试新的类型
type testField struct {
	GenericsInt     field[int]     `json:"generics_int" redis:"test_string"  default:"123" description:"测试字符串类型" category:"test"`
	GenericsInt32   field[int32]   `json:"generics_int_32" redis:"test_int_32"  default:"123" description:"测试32位整型类型" category:"test"`
	GenericsInt64   field[int64]   `json:"generics_int_64" redis:"test_int_64"  default:"123" description:"测试64位整型类型" category:"test"`
	GenericsUint    field[uint]    `json:"generics_uint" redis:"test_uint"  default:"123" description:"测试无符号整型类型" category:"test"`
	GenericsUint32  field[uint32]  `json:"generics_uint_32" redis:"test_uint_32"  default:"123" description:"测试无符号32位整型类型" category:"test"`
	GenericsUint64  field[uint64]  `json:"generics_uint_64" redis:"test_uint_64"  default:"123" description:"测试无符号64位整型类型" category:"test"`
	GenericsFloat64 field[float64] `json:"generics_float_64" redis:"test_float_64"  default:"1.23" description:"测试浮点类型" category:"test"`
	GenericsString  field[string]  `json:"generics_string" redis:"test_string"  default:"123" description:"测试字符串类型" category:"test"`
}
