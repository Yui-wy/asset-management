package data

import (
	"github.com/Yui-wy/material/app/user/service/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

// Data .
type Data struct {
	// TODO wrapped database client
	db  *gorm.DB
	log *log.Helper
}

// // new DB
// func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB{
// 	log := log.NewHelp(log.With(logger, "module", "user-service/data/gorm"))

// 	db,err := gorm.Open(mysql.Open(conf.))
// }

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}
