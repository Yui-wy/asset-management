package biz

type BaseForm struct {
	Id          int64
	AppliedAt   int64
	ApplicantId uint64
	Applicant   string
	OperatedAt  int64
	OperatorId  uint64
	Operator    string
	StateNum    int32
	State       string
	AssetId     uint64
	AssetCode   string
	AreaId      uint32
}

type BaseConfig struct {
	AreaId      []uint32
	ApplicantId uint64
	OperatorId  uint64
	StateNum    int32
	AssetId     uint64
	AssetCode   string
}
