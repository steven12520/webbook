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

func StrConvertNameByconfigID(id int)string  {

	switch id {
	case 1:
		return "鉴定估值_易鑫20张"
	case 2:
		return "便捷估值6张"
	case 4:
		return "鉴定估值_金融18张"
	case 5:
		return "便捷估值9张"
	default:
		return "无"

	}
}
func StrConvertNameByprocduct(id int)string  {

	switch id {
	case 1:
		return "常规产品"
	case 2:
		return "极速贷"
	case 3:
		return "车抵贷"
	case 4:
		return "接力贷"
	case 5:
		return "库融贷"
	case 6:
		return "交易类型"
	case 7:
		return "车抵贷(联合贷)"
	case 8:
		return "消费贷(联合贷)"
	case 9:
		return "常规产品（C）"
	case 10:
		return "【BS】常规产品（C）"
	case 11:
		return "【KG】常规产品（C）"
	case 12:
		return "【ZY】常规产品（C）"
	case 13:
		return "【KG】常规产品（TC）"
	case 14:
		return "【KG】常规产品"
	default:
		return "其他"

	}
}