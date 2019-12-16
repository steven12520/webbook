package common

func StrConvertNameBYRanges(id int)string  {

	switch id {
	case 1:
		return "缓存"
	case 2:
		return "数据库"
	case 3:
		return "缓存+数据库"
	default:
		return "其他"

	}
}