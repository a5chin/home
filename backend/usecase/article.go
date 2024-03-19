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

func (u ArticleUseCase) GetTotalViewers(ctx context.Context) (*uint, error) {
	return u.ArticleRepository.GetTotalViewers(ctx)
}

func (u ArticleUseCase) GetArticleByID(ctx context.Context, articleId string) (*entity.Article, error) {
	article, err := u.ArticleRepository.GetArticleByID(ctx, articleId)
	if err != nil {
		return nil, err
	}

	article.Viewers += 1

	if err := u.IncrementTarget(ctx, articleId, "viewers", article.Viewers); err != nil {
		return nil, err
	}

	return article, err
}
