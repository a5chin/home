package controller

import (
	"context"
	"home/entity"
)

type ArticleUseCase interface {
	GetArticles(ctx context.Context, limit, offset *int) ([]*entity.Article, error)
}

type LPUseCase interface {
	GetTotalViewers(ctx context.Context) (uint, error)
}
