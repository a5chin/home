package controller

import (
	"context"
	"home/entity"
)

type ArticleUseCase interface {
	GetArticles(ctx context.Context, limit, offset *int) ([]*entity.Article, error)
	GetArticleByID(ctx context.Context, articleId string) (*entity.Article, error)
	GetTotalViewers(ctx context.Context) (*uint, error)
}
