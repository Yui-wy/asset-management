package setting

const (
	FORM_UNKNOWN = 0
	// 用户
	FORM_SUBMITTED = 1
	FORM_CANCELED  = 2
	// 管理员
	FORM_CONFIRMED = 3
	FORM_FAIL      = 4
)

var FORM_STATE_MAP = map[int32]string{
	FORM_UNKNOWN:   "未知状态",
	FORM_SUBMITTED: "已提交",
	FORM_CANCELED:  "已取消",
	FORM_CONFIRMED: "已确认",
	FORM_FAIL:      "未通过",
}
