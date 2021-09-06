package setting

var ASSETS_STATE_MAP = map[int32]string{
	0: "未知状态, 请确认",
	1: "库内",
	2: "入库申请中",
	3: "采购中",
	4: "采购申请中",
	5: "报废",
	6: "报废申请中",
	7: "检修中",
	8: "修理中",
	9: "未知状态, 请确认",
}

const (
	SUPER_ADMIN_USER = 0
	AREA_ADMIN_USER  = 1
	AREA_USER        = 2
)

var TABLE_MAP = map[int32]string{
	AREA_ADMIN_USER: "admin_areas",
	AREA_USER:       "user_areas",
}
