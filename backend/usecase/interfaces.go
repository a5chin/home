package usecase

import (
	"context"
	"home/entity"
)

type ArticleRepository interface {
	GetArticles(ctx context.Context, limit, offset *int) ([]*entity.Article, error)
	GetArticleByID(ctx context.Context, articleId string) (*entity.Article, error)
	GetTotalViewers(ctx context.Context) (*uint, error)
	IncrementTarget(ctx context.Context, articleId string, target string, value uint) error
}
