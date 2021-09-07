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
	9: "入库失败",
}

const (
	ASSETS_STATE_UNKNOWN     = 0
	ASSETS_STATE_ST          = 1
	ASSETS_STATE_ST_APPLY    = 2
	ASSETS_STATE_ORDER       = 3
	ASSETS_STATE_ORDER_APPLY = 4
	ASSETS_STATE_SP          = 5
	ASSETS_STATE_SP_APPLY    = 6
	ASSETS_STATE_CHECKING    = 7
	ASSETS_STATE_REPAIRING   = 8
	ASSETS_STATE_ST_FAIL     = 9
)

const (
	SUPER_ADMIN_USER = 0
	AREA_ADMIN_USER  = 1
	AREA_USER        = 2
)

var TABLE_MAP = map[int32]string{
	AREA_ADMIN_USER: "admin_areas",
	AREA_USER:       "user_areas",
}
