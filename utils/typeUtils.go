package utils

import "github.com/astaxie/beego"

func TurnInt(data interface{}) int {
	switch v := data.(type) {
	case int:
		return v
	case int8:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case uintptr:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case complex64:
		beego.Error("complex64类型interface{}不能强转为int类型")
	case complex128:
		beego.Error("complex128类型interface{}不能强转为int类型")
	case string:
		beego.Error("string类型interface{}不能强转为int类型")
	case bool:
		beego.Error("bool类型interface{}不能强转为int类型")
	}
	return 0
}
