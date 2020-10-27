package service

import (
	"github.com/chinsyo/xinshou.net/controller"
	"github.com/chinsyo/xinshou.net/core"
	"github.com/chinsyo/xinshou.net/model"
	"github.com/kataras/iris/v12"

	"context"
)

type IArticleService interface {
}

var ArticleService articleService

func init() {
	ArticleService = articleService{new(controller.ArticleController)}
}

type articleService struct {
	Controller *controller.ArticleController
}

func (s *articleService) List(ctx iris.Context) {
	articles, err := s.Controller.ListArticles()
	core.SharedLogger.Infof("articleService List", articles)

	redisErr := core.RedisClient().Set(context.Background(), "name", "shawn", 0).Err()
	if redisErr != nil {
		core.SharedLogger.Infof("redisErr %s", redisErr)
	}
	val, setErr := core.RedisClient().Get(context.Background(), "name").Result()
	core.SharedLogger.Infof("val %s setErr %s", val, setErr)

	if err != nil {

	}
	ctx.JSON(articles)
}

func (s *articleService) Detail(ctx iris.Context) {
	params := ctx.Params()
	articleId, _ := params.GetUint("article_id")
	articles, err := s.Controller.ArticleDetail(articleId)
	if err != nil {

	}
	ctx.JSON(articles)
}

func (s *articleService) Create(ctx iris.Context) {

	var article model.Article
	err := ctx.ReadJSON(&article)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Article creation failure").DetailErr(err))
		return
	}

	_, err = s.Controller.CreateArticle(article)
	if err != nil {
		core.SharedLogger.Errorf("create article err: %s", err)
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Article creation failure").DetailErr(err))
		return
	}
	ctx.StatusCode(iris.StatusCreated)

}

func (s *articleService) Update(ctx iris.Context) {

}

func (s *articleService) Delete(ctx iris.Context) {

}
