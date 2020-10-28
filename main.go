package main

import (
	"github.com/chinsyo/xinshou.net/core"
	"github.com/chinsyo/xinshou.net/service"
	"github.com/kataras/iris/v12"
)

func init() {
	core.InitConfig("config.yml")
}

func main() {

	app := iris.New()
	article := app.Party("/article")
	{
		article.Get("/", service.ArticleService.List)
		article.Post("/", service.ArticleService.Create)
		article.Get("/{article_id:uint}", service.ArticleService.Detail)
		article.Delete("/{article_id:uint}", service.ArticleService.Delete)
		article.Put("/{article_id:uint}", service.ArticleService.Update)

	}

	core.SharedLogger.Infof("SharedConfig.Environment: %s", core.SharedConfig.Environment)
	app.Listen(":8080")
}
