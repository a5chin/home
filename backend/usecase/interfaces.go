package usecase

import (
	"context"
	"home/entity"
)

type ArticleRepository interface {
	GetArticles(ctx context.Context, limit, offset *int) ([]*entity.Article, error)
	GetArticleByID(ctx context.Context, articleId string) (*entity.Article, error)
}
