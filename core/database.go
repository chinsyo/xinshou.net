package core

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/chinsyo/xinshou.net/model"
)

func Conn() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres DB.name=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,                                                                              // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})
	if err != nil {

	}
	db.AutoMigrate(&model.Article{})
	SharedLogger.Infof("connect db: %s", db)
	return db
}