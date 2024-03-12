package usecase

import (
	"context"
	"home/entity"
)

type ArticleRepository interface {
	GetArticles(ctx context.Context, limit, offset *int) ([]*entity.Article, error)
}

type LPRepository interface {
	GetTotalViewers(ctx context.Context) (uint, error)
}
