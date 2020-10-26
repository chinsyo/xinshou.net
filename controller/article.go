package controller

import (
	"github.com/chinsyo/xinshou.net/model"
	"github.com/chinsyo/xinshou.net/view"
)

type ArticleController struct {
}

func (ctrl *ArticleController) ListArticles() (articles []model.Article, err error) {
	article := model.Article{Title: "Title!", Content: "Content,", ID: 1}
	articles = []model.Article{article, article}
	articles = view.ArticleView.FromArticles(articles)
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
	return article, nil
}

func (ctrl *ArticleController) UpdateArticle(article model.Article) (model.Article, error) {
	return article, nil
}

func (ctrl *ArticleController) DeleteArticle(article model.Article) (model.Article, error) {
	return article, nil
}
