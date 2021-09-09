package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/Yui-wy/asset-management/app/assets/service/internal/biz"
	"github.com/Yui-wy/asset-management/app/assets/service/internal/data"
	"github.com/Yui-wy/asset-management/pkg/setting"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestSql(t *testing.T) {
	dns := "root:root@tcp(127.0.0.1:33306)/test_assets?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed opening connection to mysql: %v", err)
	}
	if err := db.AutoMigrate(
		&data.Area{},
		&data.User{},
		&data.UserArea{},
		&data.AdminArea{},
		&data.Class{},
		&data.Asset{},
	); err != nil {
		fmt.Print(err)
	}

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", log.TraceID(),
		"span_id", log.SpanID(),
	)

	ctx := context.Background()
	dd, _, _ := data.NewData(db, logger)
	// Areas==============================================
	// ar := data.NewAreaRepo(dd, logger)
	// b, err := ar.CreateArea(ctx, &biz.Area{
	// 	AreaInfo: "Area Test3",
	// })
	// // b, err := ar.GetArea(ctx, 1)
	// // b, err := ar.GetAreasByIds(ctx, []uint32{1})
	// // b, err := ar.ListArea(ctx)
	// // b, err := ar.DeleteArea(ctx, 1)
	// // b, err := ar.UpdateArea(ctx, &biz.Area{
	// // 	Id:       2,
	// // 	AreaInfo: "Modify info2",
	// // })
	// if err != nil {
	// 	fmt.Printf("%e", err)
	// }
	// fmt.Printf("%+v", b)
	// Assets=============================================
	// assr := data.NewAssetRepo(dd, logger)
	// b, err := assr.CreatAsset(ctx, &biz.Asset{
	// 	Classes:   "030201",
	// 	AreaId:    1,
	// 	Address:   "test address",
	// 	StateNum:  1,
	// 	AppliedAt: time.Now().Unix(),
	// 	StorageAt: time.Now().Unix() + 100,
	// })
	// b, err := assr.GetAsset(ctx, 2)
	// b, err := assr.ListAssets(ctx, &biz.SearchConf{
	// 	Classes:      "030201",
	// 	LowStorageAt: 1,
	// 	// OrderBy:      "code",
	// }, 1, 5)
	// b, err := assr.UpdateAsset(ctx, &biz.Asset{
	// 	Id: 3,
	// 	// Address:   "test address 233",
	// 	StateNum: 1,
	// 	// Price:     100,
	// 	AssetInfo: "test Info~~~",
	// 	AppliedAt: time.Now().Unix(),
	// 	StorageAt: time.Now().Unix() + 100,
	// })
	// assr.DeleteAsset(ctx, 1)
	// Classes=============================================
	// cr := data.NewClassRepo(dd, logger)
	// tb := make([]*biz.Class, 0)
	// tb = append(tb, &biz.Class{
	// 	Code:    "01",
	// 	Pcode:   "",
	// 	Level:   1,
	// 	ClzInfo: "xxx1",
	// },
	// 	&biz.Class{
	// 		Code:    "12",
	// 		Pcode:   "01",
	// 		Level:   2,
	// 		ClzInfo: "xxx2",
	// 	},
	// 	&biz.Class{
	// 		Code:    "44",
	// 		Pcode:   "12",
	// 		Level:   3,
	// 		ClzInfo: "xxx44",
	// 	})
	// b, err := cr.CreateClasses(ctx, tb)
	// b, err := cr.GetClasses(ctx)
	// Users===============================================
	ur := data.NewUserRepo(dd, logger)
	b, err := ur.CreateUser(ctx, &biz.User{
		Uid:     1,
		Power:   setting.AREA_ADMIN_USER,
		AreaIds: []uint32{1, 2},
	})
	if err != nil {
		fmt.Printf("%e", err)
	}
	fmt.Printf("%+v\n", b)
	// b, err := ur.CreateUser(ctx, &biz.User{
	// 	Uid:     4,
	// 	Power:   data.AREA_ADMIN_USER,
	// 	AreaIds: []uint32{1, 2},
	// })
	// b, err = ur.CreateUser(ctx, &biz.User{
	// 	Uid:     3,
	// 	Power:   data.AREA_USER,
	// 	AreaIds: []uint32{1},
	// })
	// b, err := ur.GetUser(ctx, 3)
	// fmt.Printf("%+v\n", b)
	// b, err := ur.ListUser(ctx, data.AREA_ADMIN_USER, []uint32{1, 2})
	// fmt.Printf("%+v\n", b[0])
	// b, err := ur.UpdateUser(ctx, &biz.User{
	// 	Uid:     4,
	// 	AreaIds: []uint32{1},
	// })
	// fmt.Printf("%+v\n", b)
	// ====================================================
}
