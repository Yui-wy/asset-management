package biz

type Class struct {
	Id      uint64
	Code    string
	ClzInfo string
	Level   uint32
	Pcode   string
}

var Clzzz []*Class = []*Class{
	{
		Code:    "01",
		ClzInfo: "大类1",
		Level:   1,
	},
	{
		Code:    "02",
		ClzInfo: "大类2",
		Level:   1,
	},

	{
		Code:    "011",
		ClzInfo: "中类11",
		Level:   2,
		Pcode:   "01",
	},
	{
		Code:    "012",
		ClzInfo: "中类12",
		Level:   2,
		Pcode:   "01",
	},
	{
		Code:    "021",
		ClzInfo: "中类21",
		Level:   2,
		Pcode:   "02",
	},
	{
		Code:    "022",
		ClzInfo: "中类22",
		Level:   2,
		Pcode:   "02",
	},

	{
		Code:    "0111",
		ClzInfo: "中类111",
		Level:   2,
		Pcode:   "011",
	},
	{
		Code:    "0121",
		ClzInfo: "中类121",
		Level:   2,
		Pcode:   "012",
	},
	{
		Code:    "0211",
		ClzInfo: "中类211",
		Level:   2,
		Pcode:   "021",
	},
	{
		Code:    "0221",
		ClzInfo: "中类221",
		Level:   2,
		Pcode:   "022",
	},
}
