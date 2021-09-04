package data

import (
	"github.com/Yui-wy/asset-management/app/form/service/internal/conf"
	"github.com/Yui-wy/asset-management/pkg/util/snowflake"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewStorageRepo, NewScrappedRepo)

// Data .
type MachineCode struct {
	Datacenterid int64
	Workerid     int64
}

type Data struct {
	// TODO wrapped database client
	db  *gorm.DB
	log *log.Helper
	sf  *snowflake.Snowflake
}

// new DB
func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "form-service/data/gorm"))

	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	if err := db.AutoMigrate(&StorageForm{}, &ScrappedForm{}); err != nil {
		log.Fatal(err)
	}
	return db
}

// NewData .
func NewData(db *gorm.DB, conf *conf.Machine, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "form-service/data"))
	sf, err := snowflake.NewSnowflake(conf.Code.Datacenterid, conf.Code.Workerid)
	d := &Data{
		db:  db,
		log: log,
		sf:  sf,
	}
	cleanup := func() {
		log.Info("closing the data resources")
	}
	return d, cleanup, err
}
