package controller

import (
	"github.com/chinsyo/xinshou.net/core"
	"github.com/chinsyo/xinshou.net/model"
)

type ArticleController struct {
}

func (ctrl *ArticleController) ListArticles() (articles []model.Article, err error) {

	core.Conn().Find(&articles)
	if articles != nil {
		return
	}
	articles = []model.Article{}
	return
}

func (ctrl *ArticleController) ArticleDetail(articleId uint) (article model.Article, err error) {
	article = model.Article{Title: "Title!", Content: "Content,"}
	if articleId == article.ID {
		return
	}
	article = model.Article{Title: "Not Found!", Content: "Content,"}
	return
}

func (ctrl *ArticleController) CreateArticle(article model.Article) (model.Article, error) {
	core.Conn().Create(&article)
	return article, nil
}

func (ctrl *ArticleController) UpdateArticle(article model.Article) (model.Article, error) {
	return article, nil
}

func (ctrl *ArticleController) DeleteArticle(articleId uint) (model.Article, error) {
	var article model.Article
    if core.Conn().First(&article, articleId).RowsAffected == 0 {
        return article, core.ErrDoesNotExist
    }
    core.Conn().Delete(&article)
	return article, nil
}
