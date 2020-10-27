package controller

import (
	"github.com/chinsyo/xinshou.net/core"
	"github.com/chinsyo/xinshou.net/model"
)

type ArticleController struct {
}

func (ctrl *ArticleController) ListArticles() (articles []model.Article, err error) {
	// article := model.Article{Title: "Title!", Content: "Content,", ID: 1}
	// articles = []model.Article{article, article}
	// articles = view.ArticleView.FromArticles(articles)
	result := core.Conn().Find(&articles)
	if result.Error != nil {
		core.SharedLogger.Infof("list articles error: %s", result.Error)
	}
	return
}

func (ctrl *ArticleController) ArticleDetail(articleId uint) (article model.Article, err error) {
	article = model.Article{Title: "Title!", Content: "Content,", ID: 1}
	if articleId == article.ID {
		return
	}
	article = model.Article{Title: "Not Found!", Content: "Content,", ID: 1}
	return
}

func (ctrl *ArticleController) CreateArticle(article model.Article) (model.Article, error) {
	err := core.Conn().Create(&article).Error
	if err != nil {
		core.SharedLogger.Infof("CreateArticle error: %s", err)
	}
	return article, err
}

func (ctrl *ArticleController) UpdateArticle(article model.Article) (model.Article, error) {
	return article, nil
}

func (ctrl *ArticleController) DeleteArticle(article model.Article) (model.Article, error) {
	return article, nil
}
