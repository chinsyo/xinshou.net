package main

import (
	"github.com/chinsyo/xinshou.net/core"
	"github.com/chinsyo/xinshou.net/service"
	"github.com/kataras/iris/v12"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

func init() {
	core.InitConfig("config.yml")
}

func main() {
	// db, err := gorm.Open(postgres.New(postgres.Config{
	// 	DSN:                  "user=postgres DB.name=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai", // data source name, refer https://github.com/jackc/pgx
	// 	PreferSimpleProtocol: true,                                                                              // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	// }), &gorm.Config{})
	// if err != nil {

	// }
	// core.SharedLogger.Infof("connect db: %s", db)

	app := iris.New()
	article := app.Party("/article")
	{
		article.Get("/", service.ArticleService.List)
		article.Get("/{article_id:uint}", service.ArticleService.Detail)
		article.Post("/", service.ArticleService.Create)
	}

	core.SharedLogger.Infof("SharedConfig.Environment: %s", core.SharedConfig.Environment)
	app.Listen(":8080")
}
