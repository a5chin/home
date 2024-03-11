package usecase

import (
	"context"
	"home/entity"
)

type ArticleRepository interface {
	GetArticles(ctx context.Context, limit, offset *int) ([]*entity.Article, error)
}

type ViewerRepository interface {
	GetTotalViewers(ctx context.Context) (uint, error)
}
