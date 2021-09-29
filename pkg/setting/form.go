package setting

const (
	FORM_UNKNOWN = -1
	// 用户
	FORM_SUBMITTED = 1
	FORM_CANCELED  = 2
	// 管理员
	FORM_CONFIRMED = 3
	FORM_FAIL      = 4
)

var FORM_STATE_MAP = map[int32]string{
	FORM_UNKNOWN:   "未知状态",
	FORM_SUBMITTED: "待审批",
	FORM_CANCELED:  "已取消",
	FORM_CONFIRMED: "审批通过",
	FORM_FAIL:      "未通过",
}
