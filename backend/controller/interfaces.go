package controller

import (
	"context"
	"home/entity"
)

type ArticleUseCase interface {
	GetArticles(ctx context.Context, limit, offset *int) ([]*entity.Article, error)
}

type ViewerUseCase interface {
	GetViewers(ctx context.Context) (uint, error)
}
