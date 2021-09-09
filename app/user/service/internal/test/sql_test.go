package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/Yui-wy/asset-management/app/user/service/internal/biz"
	"github.com/Yui-wy/asset-management/app/user/service/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestSql(t *testing.T) {
	dns := "root:root@tcp(127.0.0.1:33306)/test_user?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed opening connection to mysql: %v", err)
	}
	if err := db.AutoMigrate(&data.User{}); err != nil {
		fmt.Print(err)
	}
	// ====================================================
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", log.TraceID(),
		"span_id", log.SpanID(),
	)

	dd, _, _ := data.NewData(db, logger)
	ctx := context.Background()
	str := data.NewUserRepo(dd, logger)
	// ====================================================
	u, err := str.CreateUser(ctx, &biz.User{
		Username: "test_area_admin",
		Password: "admin123",
	})
	if err != nil {
		fmt.Printf("%e", err)
	}
	fmt.Printf("%+v", u)
}
