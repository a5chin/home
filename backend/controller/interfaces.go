package controller

import (
	"context"
	"home/entity"
)

type ArticleUseCase interface {
	GetArticles(ctx context.Context, limit, offset *int) ([]*entity.Article, error)
}
