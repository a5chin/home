package usecase

import (
	"context"
	"home/entity"
)

type ArticleUseCase struct {
	ArticleRepository
}

func NewArticleUseCase(r ArticleRepository) *ArticleUseCase {
	return &ArticleUseCase{r}
}

func (u ArticleUseCase) GetArticles(ctx context.Context, limit, offset *int) ([]*entity.Article, error) {
	return u.ArticleRepository.GetArticles(ctx, limit, offset)
}
