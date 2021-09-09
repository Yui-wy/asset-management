package test

import (
	"fmt"
	"testing"

	"github.com/Yui-wy/asset-management/app/user/service/internal/data"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestSql(t *testing.T) {
	dns := "root:Mashiro&1314@tcp(127.0.0.1:3306)/test_user?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed opening connection to mysql: %v", err)
	}
	if err := db.AutoMigrate(&data.User{}); err != nil {
		fmt.Print(err)
	}
	// ====================================================
}
