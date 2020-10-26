package view

import (
	"github.com/chinsyo/xinshou.net/model"
)

const ArticleView = (articleViewImpl)(0)

type articleViewImpl int

func (v articleViewImpl) FromArticles(articles []model.Article) []model.Article {
	return articles
}
