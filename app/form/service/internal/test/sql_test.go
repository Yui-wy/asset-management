package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/Yui-wy/asset-management/app/form/service/internal/biz"
	"github.com/Yui-wy/asset-management/app/form/service/internal/conf"
	"github.com/Yui-wy/asset-management/app/form/service/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestSql(t *testing.T) {
	dns := "root:Mashiro&1314@tcp(127.0.0.1:3306)/test_form?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed opening connection to mysql: %v", err)
	}
	if err := db.AutoMigrate(&data.StorageForm{}, &data.ScrappedForm{}); err != nil {
		fmt.Print(err)
	}

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", log.TraceID(),
		"span_id", log.SpanID(),
	)

	dd, _, _ := data.NewData(db, &conf.Machine{
		Code: &conf.Machine_Code{
			Datacenterid: 0,
			Workerid:     0,
		},
	}, logger)
	ctx := context.Background()
	str := data.NewStorageRepo(dd, logger)
	// a, err := str.CreateForm(ctx, &biz.StorageForm{
	// 	BaseForm: &biz.BaseForm{
	// 		AreaId:      2,
	// 		ApplicantId: 1,
	// 		Applicant:   "testUser",
	// 		AssetId:     3,
	// 		AssetCode:   "003-030201-00002",
	// 	},
	// })
	// a, err := str.GetForm(ctx, 222138458661978112)
	// a, err := str.UpdateForm(ctx, &biz.StorageForm{
	// 	BaseForm: &biz.BaseForm{
	// 		Id:         222138458661978112,
	// 		OperatedAt: time.Now().Unix(),
	// 		OperatorId: 2,
	// 		Operator:   "testAdmin",
	// 		StateNum:   setting.FORM_CONFIRMED,
	// 	},
	// })
	// fmt.Printf("%+v\n", a.BaseForm)
	as, _, err := str.ListForm(ctx, &biz.StConfig{BaseConfig: &biz.BaseConfig{
		AreaId: []uint32{1, 2},
	}}, 1, 5)
	if err != nil {
		fmt.Printf("%e\n", err)
		return
	}
	fmt.Printf("%+v\n", as)
	for _, a := range as {
		fmt.Printf("%+v\n", a.BaseForm)
	}
	// ====================================================
}
